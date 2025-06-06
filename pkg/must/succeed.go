package must

func Succeed(err error) UntypedResult {
	return NewResult(err)
}
