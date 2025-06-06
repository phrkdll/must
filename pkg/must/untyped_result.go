package must

type UntypedResult struct {
	err error
}

type UntypedPanicer interface {
	Result
	ElsePanic()
}

func NewResult(err error) UntypedResult {
	return UntypedResult{err}
}

func (ea UntypedResult) Err() error {
	return ea.err
}

func (ea UntypedResult) ElseRespond(w ResponseWriter, statusCode int) UntypedPanicer {
	if ea.err != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(ea.err.Error()))
	}

	return ea
}

// Panics if ea.Err() is not nil
func (ea UntypedResult) ElsePanic() {
	if ea.Err() != nil {
		panic(ea.Err())
	}
}
