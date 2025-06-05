package must

type errorAction struct {
	err error
}

type responseWriter interface {
	WriteHeader(int)
	Write([]byte) (int, error)
}

func newErrorAction(err error) errorAction {
	return errorAction{err}
}

func (ea errorAction) Respond(w responseWriter, statusCode int) {
	if ea.err != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(ea.err.Error()))
	}
}

func (ea errorAction) Panic() {
	Succeed(ea.err)
}
