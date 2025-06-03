package must

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSucceed(t *testing.T) {
	type testCase struct {
		name string
		err  error
	}

	testCases := []testCase{
		{
			name: "function returns error",
			err:  errors.New("something was wrong"),
		},
		{
			name: "function returns nil",
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err == nil {
				assert.NotPanics(t, func() { Succeed(tc.err) })
			} else {
				assert.Panics(t, func() { Succeed(tc.err) })
			}
		})
	}
}

func TestReturn(t *testing.T) {
	type testCase struct {
		name string
		val  any
		err  error
	}

	testCases := []testCase{
		{
			name: "error is not nil",
			val:  nil,
			err:  errors.New("something was wrong"),
		},
		{
			name: "value is string",
			val:  "test",
			err:  nil,
		},
		{
			name: "value is float",
			val:  1.1,
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err == nil {
				res := Return(tc.val, tc.err)

				assert.Equal(t, tc.val, res)
			} else {
				assert.Panics(t, func() { Return(tc.val, tc.err) })
			}
		})
	}
}

func TestReturnOne(t *testing.T) {
	type testCase struct {
		name string
		val  any
		err  error
	}

	testCases := []testCase{
		{
			name: "error is not nil",
			val:  nil,
			err:  errors.New("something was wrong"),
		},
		{
			name: "value is string",
			val:  "test",
			err:  nil,
		},
		{
			name: "value is float",
			val:  1.1,
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err == nil {
				res := ReturnOne(tc.val, tc.err)

				assert.Equal(t, tc.val, res)
			} else {
				assert.Panics(t, func() { ReturnOne(tc.val, tc.err) })
			}
		})
	}
}

func TestReturnTwo(t *testing.T) {
	type testCase struct {
		name string
		vals []any
		err  error
	}

	testCases := []testCase{
		{
			name: "error is not nil",
			vals: nil,
			err:  errors.New("something was wrong"),
		},
		{
			name: "value is string",
			vals: []any{"test1", "test2"},
			err:  nil,
		},
		{
			name: "value is float",
			vals: []any{1.1, 2.2},
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err == nil {
				res0, res1 := ReturnTwo(tc.vals[0], tc.vals[1], tc.err)

				assert.Equal(t, tc.vals[0], res0)
				assert.Equal(t, tc.vals[1], res1)
			} else {
				assert.Panics(t, func() { ReturnTwo(tc.vals[0], tc.vals[1], tc.err) })
			}
		})
	}
}

func TestReturnThree(t *testing.T) {
	type testCase struct {
		name string
		vals []any
		err  error
	}

	testCases := []testCase{
		{
			name: "error is not nil",
			vals: nil,
			err:  errors.New("something was wrong"),
		},
		{
			name: "value is string",
			vals: []any{"test1", "test2", "test3"},
			err:  nil,
		},
		{
			name: "value is float",
			vals: []any{1.1, 2.2, 3.3},
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err == nil {
				res0, res1, res2 := ReturnThree(tc.vals[0], tc.vals[1], tc.vals[2], tc.err)

				assert.Equal(t, tc.vals[0], res0)
				assert.Equal(t, tc.vals[1], res1)
				assert.Equal(t, tc.vals[2], res2)
			} else {
				assert.Panics(t, func() { ReturnThree(tc.vals[0], tc.vals[1], tc.vals[2], tc.err) })
			}
		})
	}
}
