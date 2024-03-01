module github.com/wokill/goc

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/go-github v17.0.0+incompatible
	github.com/hashicorp/go-retryablehttp v0.6.6
	github.com/julienschmidt/httprouter v1.2.0
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/olekukonko/tablewriter v0.0.4
	github.com/qiniu/api.v7/v7 v7.5.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tongjingran/copy v1.4.2
	golang.org/x/mod v0.3.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/tools v0.0.0-20200730221956-1ac65761fe2c
	k8s.io/test-infra v0.0.0-20200511080351-8ac9dbfab055
)

replace github.com/mitchellh/osext v0.0.0-20151018003038-5e2d6d41470f => github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0

exclude github.com/mitchellh/osext v0.0.0-20151018003038-5e2d6d41470f
