package cover

import (
	"fmt"
	"net/url"
)

func (c *client) RedirectService(address, name, port, mod string) ([]byte, error) {
	parse, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	if len(port) == 0 {
		port = parse.Port()
	}
	_, res, err := c.do("GET", fmt.Sprintf("http://%s:%s/v1/cover/html?service=%s&mod=%s", parse.Hostname(), port, name, mod), "", nil)
	return res, err
}
