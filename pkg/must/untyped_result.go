package must

import "encoding/json"

type UntypedResult struct {
	err     error  `json:"-"`
	Success bool   `json:"success"`
	Error   bool   `json:"error"`
	Message string `json:"message,omitempty"`
}

type ResultAction func()

func NewResult(err error) UntypedResult {
	var errorMessage string
	hasError := err != nil
	if hasError {
		errorMessage = err.Error()
	}

	return UntypedResult{err, !hasError, hasError, errorMessage}
}

func (res UntypedResult) ElseRespond(w ResponseWriter, statusCode int) {
	if res.err != nil {
		json, err := json.Marshal(&res)
		if err != nil {
			panic(err)
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(json)

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
