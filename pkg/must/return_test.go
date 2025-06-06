package must_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/phrkdll/must/pkg/must"
	"github.com/stretchr/testify/assert"
)

func TestReturn(t *testing.T) {
	type testCase struct {
		name       string
		val        any
		err        error
		statusCode int
	}

	testCases := []testCase{
		{
			name:       "error is not nil",
			val:        nil,
			err:        errors.New("something went wrong"),
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "value is string",
			val:        "test",
			err:        nil,
			statusCode: http.StatusOK,
		},
		{
			name:       "value is float",
			val:        1.1,
			err:        nil,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var writer must.MockResponseWriter
			panicer := must.Return(tc.val, tc.err).ElseRespond(&writer, tc.statusCode)

			if tc.err != nil {
				assert.Equal(t, tc.err, writer.Error)
				assert.Equal(t, tc.err, panicer.Err())
				assert.Equal(t, tc.statusCode, writer.StatusCode)
				assert.Panics(t, func() { panicer.ElsePanic() })
			} else {
				assert.NotPanics(t, func() { panicer.ElsePanic() })
				assert.Equal(t, tc.val, panicer.ElsePanic())
				assert.Nil(t, panicer.Err())
				assert.Nil(t, writer.Error)
				assert.Zero(t, writer.StatusCode)
			}
		})
	}
}
