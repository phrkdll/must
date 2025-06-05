package must

func SucceedOr(err error) errorAction {
	return newErrorAction(err)
}
