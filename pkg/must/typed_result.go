package must

type TypedResult[T any] struct {
	UntypedResult
	val T
}

type TypedPanicer[T any] interface {
	Result
	ElsePanic() T
}

func NewTypedResult[T any](ea UntypedResult, val T) TypedResult[T] {
	return TypedResult[T]{ea, val}
}

func (ea TypedResult[T]) ElseRespond(w ResponseWriter, statusCode int) TypedPanicer[T] {
	if ea.Err() != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(ea.Err().Error()))
	}

	return ea
}

// Panics if ea.Err() is not nil
// Otherwise returns the value from the original function
func (ea TypedResult[T]) ElsePanic() T {
	if ea.Err() != nil {
		panic(ea.Err())
	}

	return ea.val
}
