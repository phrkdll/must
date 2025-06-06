package must

func Return[T any](val T, err error) TypedResult[T] {
	return NewTypedResult(
		Succeed(err),
		val,
	)
}
