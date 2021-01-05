package fallback

import "net/http"

type interceptor struct {
	http.ResponseWriter
	isNotFound bool
}

func (this *interceptor) WriteHeader(status int) {
	if status != http.StatusNotFound {
		this.ResponseWriter.WriteHeader(status)
		return
	}

	this.isNotFound = true
}

func (this *interceptor) Write(p []byte) (int, error) {
	if this.isNotFound {
		return len(p), nil
	}

	return this.ResponseWriter.Write(p)
}

func Fallback(fallback http.Handler, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		interceptor := &interceptor{ResponseWriter: rw}
		handler.ServeHTTP(interceptor, r)

		if interceptor.isNotFound {
			fallback.ServeHTTP(rw, r)
		}
	})
}
