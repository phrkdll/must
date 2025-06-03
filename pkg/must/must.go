package must

func Succeed(err error) {
	if err != nil {
		panic(err)
	}
}

func Return[T any](val T, err error) T {
	Succeed(err)

	return val
}

func ReturnOne[T any](val T, err error) T {
	return Return(val, err)
}

func ReturnTwo[T1, T2 any](val1 T1, val2 T2, err error) (T1, T2) {
	Succeed(err)

	return val1, val2
}

func ReturnThree[T1, T2, T3 any](val1 T1, val2 T2, val3 T3, err error) (T1, T2, T3) {
	Succeed(err)

	return val1, val2, val3
}
