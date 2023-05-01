module github.com/pjvds/tunl

go 1.20

replace github.com/inconshreveable/go-vhost => github.com/pjvds/go-vhost v0.0.0-20201229150248-206eee94f4aa

replace github.com/armon/go-metrics => ./go-metrics

// replace github.com/armon/go-metrics => github.com/hashicorp/go-metrics v0.4.2-0.20221220172610-8cabd9eab1be

replace github.com/hashicorp/golang-lru => github.com/hashicorp/golang-lru/v2 v2.0.2

replace github.com/hashicorp/go-immutable-radix => github.com/hashicorp/go-immutable-radix/v2 v2.0.0

require (
	github.com/Masterminds/semver v1.5.0
	github.com/armon/go-metrics v0.4.2-0.20221220172610-8cabd9eab1be
	github.com/docker/docker v23.0.5+incompatible
	github.com/getspine/go-metrics-honeycomb v0.0.0-20180125230210-7f1304f1bc83
	github.com/gobwas/glob v0.2.3
	github.com/goji/httpauth v0.0.0-20160601135302-2da839ab0f4d
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/sessions v1.2.1
	github.com/hashicorp/yamux v0.1.1
	github.com/inconshreveable/go-vhost v1.0.0
	github.com/mdp/qrterminal/v3 v3.0.0
	github.com/pkg/errors v0.9.1
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	github.com/rs/xid v1.5.0
	github.com/urfave/cli/v2 v2.25.1
	github.com/yelinaung/go-haikunator v0.0.0-20221222235932-36bf4c441150
	go.uber.org/zap v1.24.0
	golang.org/x/net v0.6.0
)

require (
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/atotto/clipboard v0.1.4
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/facebookgo/limitgroup v0.0.0-20150612190941-6abd8d71ec01 // indirect
	github.com/facebookgo/muster v0.0.0-20150708232844-fd3d7953fd52 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/hashicorp/go-immutable-radix/v2 v2.0.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.0 // indirect
	github.com/honeycombio/libhoney-go v1.18.0 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/moby/term v0.0.0-20230430220526-1849d9c42740 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	gopkg.in/alexcesaro/statsd.v2 v2.0.0 // indirect
	gotest.tools/v3 v3.4.0 // indirect
	rsc.io/qr v0.2.0 // indirect
)
