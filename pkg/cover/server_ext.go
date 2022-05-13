package cover

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/tools/cover"
	"k8s.io/test-infra/gopherage/pkg/cov"
)

func (s *server) SetRedirectPort(port string) {
	s.RedirectPort = port
}

// RedirectParam is redirect url
type RedirectParam struct {
	Service  string `json:"service" form:"service" binding:"required"`
	Port     string `json:"port" form:"port"`
	Mode     string `json:"mode" form:"mode"`         // 显示默认 func 函数类型，默认代码 code
	Download bool   `json:"download" form:"download"` // true 下载模式
}

// redirect other url profile
func (s *server) redirect(c *gin.Context) {
	var service RedirectParam
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adders := s.Store.Get(service.Service)

	if len(adders) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "service is empty or service is not register"})
		return
	}
	if len(service.Port) == 0 {
		service.Port = s.RedirectPort
	}
	resp, err := NewWorker(adders[0]).RedirectService(adders[0], service.Service, service.Port, service.Mode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if service.Download {
		suffix := ".html"
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-disposition", "attachment; filename=\""+service.Service+"\"_代码覆盖率"+suffix)
		c.String(http.StatusOK, string(resp))
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, string(resp))
}

func (s *server) profileContent(body ProfileParam) ([]*cover.Profile, error) {
	allInfos := s.Store.GetAll()
	filterAddrList, err := filterAddrs(body.Service, body.Address, body.Force, allInfos)
	if err != nil {
		return nil, err
	}

	var mergedProfiles = make([][]*cover.Profile, 0)
	for _, addr := range filterAddrList {
		pp, err := NewWorker(addr).Profile(ProfileParam{})
		if err != nil {
			if body.Force {
				log.Warnf("get profile from [%s] failed, error: %s", addr, err.Error())
				continue
			}
			return nil, fmt.Errorf("failed to get profile from %s, error %s", addr, err.Error())
		}

		profile, err := convertProfile(pp)
		if err != nil {
			return nil, err
		}
		mergedProfiles = append(mergedProfiles, profile)
	}

	if len(mergedProfiles) == 0 {
		return nil, errors.New("no profiles")
	}

	merged, err := cov.MergeMultipleProfiles(mergedProfiles)
	if err != nil {
		return nil, err
	}

	if len(body.CoverFilePatterns) > 0 {
		merged, err = filterProfile(body.CoverFilePatterns, merged)
		if err != nil {
			return nil, err
		}
	}

	if len(body.SkipFilePatterns) > 0 {
		merged, err = skipProfile(body.SkipFilePatterns, merged)
		if err != nil {
			return nil, err
		}
	}
	return merged, nil
}
