package must_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/phrkdll/must/pkg/must"
	"github.com/stretchr/testify/assert"
)

func TestSucceed(t *testing.T) {
	type testCase struct {
		name       string
		err        error
		statusCode int
	}

	testCases := []testCase{
		{
			name:       "function returns error",
			err:        errors.New("something went wrong"),
			statusCode: http.StatusBadRequest,
		},
		{
			name: "function returns nil",
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var writer must.MockResponseWriter
			panicer := must.Succeed(tc.err).ElseRespond(&writer, tc.statusCode)

			if tc.err != nil {
				assert.Equal(t, tc.err, writer.Error)
				assert.Equal(t, tc.err, panicer.Err())
				assert.Equal(t, tc.statusCode, writer.StatusCode)
				assert.Panics(t, func() { must.Succeed(tc.err).ElsePanic() })
			} else {
				assert.NotPanics(t, func() { must.Succeed(tc.err).ElsePanic() })
				assert.Nil(t, panicer.Err())
				assert.Nil(t, writer.Error)
				assert.Zero(t, writer.StatusCode)
			}
		})
	}
}
