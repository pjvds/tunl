module github.com/pjvds/tunl

go 1.15

require (
	github.com/foomo/simplecert v1.8.1 // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/yamux v0.0.0-20200609203250-aecfd211c9ce
	github.com/inconshreveable/go-vhost v0.0.0-20160627193104-06d84117953b
	github.com/pkg/errors v0.9.1
	github.com/urfave/cli v1.22.5
	github.com/urfave/cli/v2 v2.3.0
	github.com/yelinaung/go-haikunator v0.0.0-20150320004105-1249cae259af
	go.uber.org/zap v1.16.0
)

replace github.com/inconshreveable/go-vhost => github.com/pjvds/go-vhost v0.0.0-20201229150248-206eee94f4aa
