package must_test

import (
	"errors"
	"testing"

	"github.com/phrkdll/must/pkg/must"
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
			err:  errors.New("something went wrong"),
		},
		{
			name: "function returns nil",
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err == nil {
				assert.NotPanics(t, func() { must.Succeed(tc.err) })
			} else {
				assert.Panics(t, func() { must.Succeed(tc.err) })
			}
		})
	}
}
