package must

func Succeed(err error) {
	if err != nil {
		panic(err)
	}
}
