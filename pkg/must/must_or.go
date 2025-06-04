package must

type errorAction struct {
	err error
}

type ResponseWriter interface {
	WriteHeader(int)
	Write([]byte) (int, error)
}

func newErrorAction(err error) errorAction {
	return errorAction{err: err}
}

func (ea errorAction) Respond(w ResponseWriter, status int) {
	if ea.err != nil {
		w.WriteHeader(status)
		w.Write([]byte(ea.err.Error()))
	}
}

func (ea errorAction) Panic() {
	Succeed(ea.err)
}

func SucceedOr(err error) errorAction {
	return newErrorAction(err)
}
