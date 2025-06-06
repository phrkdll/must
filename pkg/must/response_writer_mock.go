package must

import "errors"

type MockResponseWriter struct {
	StatusCode int
	Error      error
}

func (rw *MockResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
}

func (rw *MockResponseWriter) Write(body []byte) (int, error) {
	rw.Error = errors.New(string(body))

	return 0, nil
}
