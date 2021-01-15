package templates

import (
	_ "embed"
	"html/template"
	"io"
)

//go:embed http-client-error.tmpl
var httpClientError string

type HttpClientErrorInput struct {
	RemoteAddress     string
	LocalHostname     string
	LocalAddress      string
	TunlClientVersion string
	ErrMessage        string
	ErrType           string
	ErrDetails        string
	Year              int
}

var httpClientErrorTemplate *template.Template

func init() {
	httpClientErrorTemplate = template.Must(template.New("").Parse(httpClientError))
}

func HttpClientError(writer io.Writer, input HttpClientErrorInput) error {

	return httpClientErrorTemplate.ExecuteTemplate(writer, "http-client-error", input)
}
