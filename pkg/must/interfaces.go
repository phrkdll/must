package must

type Result interface {
	Err() error
}

type ResponseWriter interface {
	WriteHeader(int)
	Write([]byte) (int, error)
}

type Responder interface {
	Result
	ElseRespond(ResponseWriter, int)
}
