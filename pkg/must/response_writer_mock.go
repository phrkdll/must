package must

import "net/http"

type MockResponseWriter struct {
	StatusCode int
	Error      *string
}

func (rw *MockResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
}

func (rw *MockResponseWriter) Write(body []byte) (int, error) {
	err := string(body)
	rw.Error = &err

	return 0, nil
}

func (rw *MockResponseWriter) Header() http.Header {
	return http.Header{}
}
