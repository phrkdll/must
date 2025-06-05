package must

type function func(error)

func Recover() {
	_ = recover()
}

func RecoverWith(fn function) {
	err := recover()

	if err != nil {
		fn(err.(error))
	}
}
