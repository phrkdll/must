package must

import (
	"encoding/json"
)

type TypedResult[T any] struct {
	UntypedResult
	Val T `json:"val"`
}

type TypedResultAction[T any] func(T)

func NewTypedResult[T any](res UntypedResult, val T) TypedResult[T] {
	return TypedResult[T]{res, val}
}

func (res TypedResult[T]) ElseRespond(w ResponseWriter, statusCode int) T {
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

	return res.Val
}

func (res TypedResult[T]) ElseExecute(fn TypedResultAction[T]) T {
	if res.err != nil {
		fn(res.Val)

		panic(res.err)
	}

	return res.Val
}

func (res TypedResult[T]) ElsePanic() T {
	if res.err != nil {
		panic(res.err)
	}

	return res.Val
}
