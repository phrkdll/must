package must

func Return[T any](data T, err error) TypedResult[T] {
	return NewTypedResult(
		Succeed(err),
		data,
	)
}
