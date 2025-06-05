package must_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/phrkdll/must/pkg/must"
	"github.com/stretchr/testify/assert"
)

type mockResponseWriter struct {
	StatusCode int
	Error      error
}

func (rw *mockResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
}

func (rw *mockResponseWriter) Write(body []byte) (int, error) {
	rw.Error = errors.New(string(body))

	return 0, nil
}

func TestSucceedOr(t *testing.T) {
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
			var mrw mockResponseWriter
			must.SucceedOr(tc.err).Respond(&mrw, tc.statusCode)

			if tc.err != nil {
				assert.Equal(t, tc.err, mrw.Error)
				assert.Equal(t, tc.statusCode, mrw.StatusCode)
				assert.Panics(t, func() { must.SucceedOr(tc.err).Panic() })
			} else {
				assert.NotPanics(t, func() { must.SucceedOr(tc.err).Panic() })
				assert.Nil(t, mrw.Error)
				assert.Zero(t, mrw.StatusCode)
			}
		})
	}
}
