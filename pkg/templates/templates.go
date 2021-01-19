package templates

import (
	_ "embed"
	"html/template"
	"io"
)

//go:embed http-client-error.tmpl
var httpClientError string

//go:embed password.tmpl
var password string

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
var passwordTemplate *template.Template

func init() {
	httpClientErrorTemplate = template.Must(template.New("http-client-error").Parse(httpClientError))

	passwordTemplate = template.Must(template.New("password").Parse(password))
}

type PasswordInput struct {
	Message string
}

func Password(writer io.Writer, input PasswordInput) error {
	return passwordTemplate.Execute(writer, input)
}

func HttpClientError(writer io.Writer, input HttpClientErrorInput) error {
	return httpClientErrorTemplate.Execute(writer, input)
}
