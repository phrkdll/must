package must

type Result interface {
	Err() error
}

type ResponseWriter interface {
	Header() Header
	WriteHeader(int)
	Write([]byte) (int, error)
}

type Header interface {
	Add(string, string)
}

type Responder interface {
	Result
	ElseRespond(ResponseWriter, int)
}
