package must

type ExecuteTester struct {
	Called bool
}

func (t *ExecuteTester) TestUntyped() {
	t.Called = true
}

func (t *ExecuteTester) TestTyped(val any) {
	t.Called = true
}
