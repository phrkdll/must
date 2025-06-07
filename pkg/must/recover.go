package must

type RecoverAction func(error)

// Recover without custom error handling
// (ignores the error returned by Go's recover function)
func Recover() {
	_ = recover()
}

// Recover with custom error handling
// (passes the error returned by Go's recover function to the passed RecoverAction)
func RecoverWith(fn RecoverAction) {
	err := recover()

	if err != nil {
		fn(err.(error))
	}
}
