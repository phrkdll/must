package must

type recoverFunction func(error)

// Recover without custom error handling
// (ignores the error returned by Go's recover function)
func Recover() {
	_ = recover()
}

// Recover with custom error handling
// (passes the error returned by Go's recover function to the passed recoverFunction)
func RecoverWith(fn recoverFunction) {
	err := recover()

	if err != nil {
		fn(err.(error))
	}
}
