package must

import "net/http"

type Result interface {
	Err() error
}

type ResponseWriter interface {
	Header() http.Header
	WriteHeader(int)
	Write([]byte) (int, error)
}

type Responder interface {
	Result
	ElseRespond(ResponseWriter, int)
}
