package commands

import (
	"crypto/tls"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/pjvds/tunl/pkg/tunnel"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var downstreamErrTemplate, _ = template.New("").Parse(`
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="icon" href="/docs/4.0/assets/img/favicons/favicon.ico">

    <title>tunl client proxy error</title>

		<script src="https://kit.fontawesome.com/f95edda927.js" crossorigin="anonymous"></script>

    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-giJF6kkoqNQ00vy+HMDP7azOuL0xtbfIcaT9wjKHr8RbDVddVHyTfAAsrekwKmP1" crossorigin="anonymous">

		<style>
		html,
		body {
			height: 100%;
		}

		body {
			display: -ms-flexbox;
			display: -webkit-box;
			display: flex;
			-ms-flex-align: center;
			-ms-flex-pack: center;
			-webkit-box-align: center;
			align-items: center;
			-webkit-box-pack: center;
			justify-content: center;
			padding-top: 40px;
			padding-bottom: 40px;
			background-color: #f5f5f5;
		}

		.list-group{
			margin-left: 0em;
		}

		.form-signin {
			width: 100%;
			max-width: 500px;
			padding: 15px;
			margin: 0 auto;
		}
		.form-signin .checkbox {
			font-weight: 400;
		}
		.form-signin .form-control {
			position: relative;
			box-sizing: border-box;
			height: auto;
			padding: 10px;
			font-size: 16px;
		}
		.form-signin .form-control:focus {
			z-index: 2;
		}
		.form-signin input[type="email"] {
			margin-bottom: -1px;
			border-bottom-right-radius: 0;
			border-bottom-left-radius: 0;
		}
		.form-signin input[type="password"] {
			margin-bottom: 10px;
			border-top-left-radius: 0;
			border-top-right-radius: 0;
		}
		</style>
  </head>

  <body >
    <div class="form-signin">
      <h1 class="h2 mb-3 font-weight-normal">Oops! 
			<br/>
			{{.ErrMessage}}</h1>
			<p>
				The public address <code>{{.RemoteAddress}}</code> received a request and forwared it succesfully to the tunl client running at <code>{{.LocalHostname}}</code>. But this client failed to proxy the request to http target <code>{{.LocalAddress}}</code>.
			</p>
			<p>
			Please make sure the target HTTP service is running at <a href="{{.LocalAddress}}">{{.LocalAddress}}</a> and can be reached from the local network.
			</p>
      <h1 class="h2 mb-3 font-weight-normal">Details</h1>
			<p>
			<div class="alert alert-secondary" role="alert">
			<pre>[{{.ErrType}}]
{{.ErrDetails}}</pre>
			</div>
			</p>

      <h1 class="h2 mb-3 font-weight-normal">Status</h1>
			<p>
			<ul class="fa-ul list-group">
			<li class="list-group-item list-group-item-success"><span class="fa-li"><i class="fas fa-check"></i></span>Public address received the request</li>
			<li class="list-group-item list-group-item-success"><span class="fa-li"><i class="fas fa-check"></i></span>Request forward to local client</code></li>
			<li class="list-group-item list-group-item-success"><span class="fa-li"><i class="fas fa-check"></i></span>Local client received the request</li>
			<li class="list-group-item list-group-item-danger"><span class="fa-li"><i class="fas fa-times"></i></span>Local client failed to proxy request to target url: <br /><code>{{.ErrMessage}}</code></li>
			</ul>
			</p>
      <button class="btn btn-lg btn-primary btn-block" onClick="window.location.reload();">Retry</button>
      <p class="mt-5 mb-3 text-muted">&copy; tunl.es {{.Year}}</p>
    </div>
  </body>
</html>
`)

var HttpCommand = &cli.Command{
	Name: "http",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "access-log",
			Usage: "Print http requests in Apache Log format to stderr",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "insecure",
			Usage: "Skip TLS verification for local address (this does not effect TLS between the tunl client and server or the public address)",
			Value: true,
		},
	},
	ArgsUsage: "<url>",
	Usage:     "Expose a HTTP service via a public address",
	Action: func(ctx *cli.Context) error {
		var targetURL *url.URL
		target := ctx.Args().First()
		if len(target) == 0 {
			fmt.Fprint(os.Stderr, "You must specify the <url> argument\n\n")
			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
		}

		if !strings.Contains(target, "://") {
			target = "http://" + target
		}

		parsed, err := url.Parse(target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid <url> argument value: %v\n\n", err)
			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
		}
		targetURL = parsed

		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		if ctx.Bool("insecure") {
			proxy.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}

		originalDirector := proxy.Director

		proxy.Director = func(request *http.Request) {
			originalDirector(request)
			request.Host = targetURL.Host
		}

		host := ctx.String("host")
		if len(host) == 0 {
			fmt.Fprint(os.Stderr, "Host cannot be empty\nSee --host flag for more information.\n\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return cli.Exit("Host cannot be empty.", 1)
		}

		hostURL, err := url.Parse(host)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Host value invalid: %v\nSee --host flag for more information.\n\n", err)

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return nil
		}

		hostnameWithoutPort := hostURL.Hostname()
		if len(hostnameWithoutPort) == 0 {
			fmt.Fprintf(os.Stderr, "Host hostname cannot be empty, see --host flag for more information.\n\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return nil
		}

		tunnel, err := tunnel.OpenHTTP(ctx.Context, zap.NewNop(), hostURL)
		if err != nil {
			return cli.Exit(err.Error(), 18)
		}

		PrintTunnel(tunnel.Address(), target)

		handler := handlers.LoggingHandler(os.Stdout, proxy)

		proxy.ErrorHandler = func(response http.ResponseWriter, request *http.Request, err error) {
			hostname, _ := os.Hostname()
			if len(hostname) == 0 {
				hostname = "<unknown>"
			}

			fmt.Println(err)

			var unwrapped = err

			for next := errors.Unwrap(unwrapped); next != nil; next = errors.Unwrap(unwrapped) {
				unwrapped = next
			}

			response.WriteHeader(http.StatusBadGateway)

			downstreamErrTemplate.Execute(response, struct {
				RemoteAddress     string
				LocalHostname     string
				LocalAddress      string
				TunlClientVersion string
				ErrMessage        string
				ErrType           string
				ErrDetails        string
				Year              int
			}{
				RemoteAddress:     tunnel.Address(),
				LocalHostname:     hostname,
				LocalAddress:      target,
				TunlClientVersion: ctx.App.Version,
				ErrMessage:        unwrapped.Error(),
				ErrType:           fmt.Sprintf("%T", err),
				ErrDetails:        err.Error(),
				Year:              time.Now().Year(),
			})
		}

		if err := http.Serve(tunnel, handler); err != nil {
			return cli.Exit("fatal error: "+err.Error(), 1)
		}

		return nil
	},
}
