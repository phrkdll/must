package must

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

func (rw *MockResponseWriter) Header() Header {
	return &MockHeader{}
}

type MockHeader struct {
	headers map[string]string
}

func (h *MockHeader) Add(key, value string) {
	if h.headers == nil {
		h.headers = make(map[string]string)
	}
	h.headers[key] = value
}
