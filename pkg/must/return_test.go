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
		data       any
		err        error
		statusCode int
	}

	testCases := []testCase{
		{
			name:       "error is not nil",
			data:       nil,
			err:        errors.New("something went wrong"),
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "value is string",
			data:       "test",
			err:        nil,
			statusCode: http.StatusOK,
		},
		{
			name:       "value is float",
			data:       1.1,
			err:        nil,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var writer must.MockResponseWriter
			tester := &must.ExecuteTester{}

			if tc.err != nil {
				assert.Panics(t, func() { must.Return(tc.data, tc.err).ElseRespond(&writer, tc.statusCode) })
				assert.Panics(t, func() { must.Return(tc.data, tc.err).ElseExecute(tester.TestTyped) })
				assert.Panics(t, func() { must.Return(tc.data, tc.err).ElsePanic() })
				assert.True(t, tester.Called)
				assert.Contains(t, *writer.Error, tc.err.Error())
				assert.Equal(t, tc.statusCode, writer.StatusCode)
			} else {
				var result any
				assert.NotPanics(t, func() { result = must.Return(tc.data, tc.err).ElseRespond(&writer, tc.statusCode) })
				assert.NotPanics(t, func() { must.Return(tc.data, tc.err).ElseExecute(tester.TestTyped) })
				assert.NotPanics(t, func() { must.Return(tc.data, tc.err).ElsePanic() })
				assert.False(t, tester.Called)
				assert.Equal(t, tc.data, result)
				assert.Nil(t, writer.Error)
				assert.Zero(t, writer.StatusCode)
			}
		})
	}
}
