package must

type action interface {
	Err() error
}

type responder interface {
	action
	Respond(responseWriter, int)
}

type panicer interface {
	action
	Panic()
}

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

func (ea errorAction) Err() error {
	return ea.err
}

func (ea errorAction) Respond(w responseWriter, statusCode int) panicer {
	if ea.err != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(ea.err.Error()))
	}

	return ea
}

func (ea errorAction) Panic() {
	Succeed(ea.err)
}
