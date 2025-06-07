package must

type UntypedResult struct {
	err error
}

type ResultAction func()

func NewResult(err error) UntypedResult {
	return UntypedResult{err}
}

func (res UntypedResult) ElseRespond(w ResponseWriter, statusCode int) {
	if res.err != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(res.err.Error()))

		panic(res.err)
	}
}

func (res UntypedResult) ElseExecute(fn ResultAction) {
	if res.err != nil {
		fn()

		panic(res.err)
	}
}

func (res UntypedResult) ElsePanic() {
	if res.err != nil {
		panic(res.err)
	}
}
