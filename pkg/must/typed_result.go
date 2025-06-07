package must

type TypedResult[T any] struct {
	UntypedResult
	val T
}

type TypedResultAction[T any] func(T)

func NewTypedResult[T any](res UntypedResult, val T) TypedResult[T] {
	return TypedResult[T]{res, val}
}

func (res TypedResult[T]) ElseRespond(w ResponseWriter, statusCode int) T {
	if res.err != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(res.err.Error()))

		panic(res.err)
	}

	return res.val
}

func (res TypedResult[T]) ElseExecute(fn TypedResultAction[T]) T {
	if res.err != nil {
		fn(res.val)

		panic(res.err)
	}

	return res.val
}

func (res TypedResult[T]) ElsePanic() T {
	if res.err != nil {
		panic(res.err)
	}

	return res.val
}
