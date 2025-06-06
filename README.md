# must ![coverage](https://raw.githubusercontent.com/phrkdll/must/badges/.badges/main/coverage.svg)

A simple package providing easy access to several functions that cover the must-pattern, which is often utilized in Go.

## Usage

Get the dependency:

```shell
go get github.com/phrkdll/must
```

Use the package:

```go
package main

import (
    "net/http"
    "github.com/phrkdll/must/pkg/must"
)

func divide(dividend, divisor float64) (float64, error) {
    if divisor == 0.0 {
        return 0.0, errors.New("can't divide by zero")
    }

    return dividend / divisor, nil
}

func someOtherFunc(val string) error {
    if val == "" {
        return errors.New("val must not be empty")
    }

    return nil
}

func someHandler(w http.ResponseWriter, r *http.Request) {
	defer must.Recover()
	// OR
	defer must.RecoverWith(func(err error) { 
		// ... 
	})

	must.Succeed(someOtherFunc("")).ElseRespond(w, http.StatusInternalServerError).ElsePanic()
	// OR
	quotient := must.Return(divide(1, 2)).ElseRespond(w, http.StatusInternalServerError).ElsePanic()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte{byte(quotient)})
}
```
