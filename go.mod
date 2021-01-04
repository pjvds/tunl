module github.com/pjvds/tunl

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/yamux v0.0.0-20200609203250-aecfd211c9ce
	github.com/inconshreveable/go-vhost v0.0.0-20160627193104-06d84117953b
	github.com/kr/pretty v0.2.1 // indirect
	github.com/pjvds/backoff v0.0.0-20151029185359-615bd1fa5d8a
	github.com/pkg/errors v0.9.1
	github.com/rs/xid v1.2.1
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/urfave/cli/v2 v2.3.0
	github.com/yelinaung/go-haikunator v0.0.0-20150320004105-1249cae259af
	go.uber.org/zap v1.16.0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/mod v0.4.0 // indirect
	golang.org/x/tools v0.0.0-20201202200335-bef1c476418a // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)

replace github.com/inconshreveable/go-vhost => github.com/pjvds/go-vhost v0.0.0-20201229150248-206eee94f4aa
