package templates

import (
	"html/template"
	"io"
	"io/ioutil"

	"github.com/markbates/pkger"
)

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

var httpClientError *template.Template

func HttpClientError(writer io.Writer, input HttpClientErrorInput) error {
	return httpClientError.Execute(writer, input)
}

func init() {
	httpClientErrorFile, err := pkger.Open("/assets/templates/http-client-error.tmpl")
	if err != nil {
		panic(err)
	}

	httpClientErrorContent, err := ioutil.ReadAll(httpClientErrorFile)
	if err != nil {
		panic(err)
	}

	httpClientError, err = template.New("http-client-error").Parse(string(httpClientErrorContent))
	if err != nil {
		panic(err)
	}
}
