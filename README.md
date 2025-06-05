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
    "github.com/phrkdll/must/pkg/must"
)

func someFunc(val string) error {
    if val == "" {
        return errors.New("val must not be empty")
    }

    return nil
}
    
must.Succeed(someFunc(""))
// Alternative usage
must.SucceedOr(someFunc("")).Panic()

func divide(dividend, divisor float64) (float64, error) {
    if divisor == 0.0 {
        return 0.0, errors.New("can't divide by zero")
    }

    return dividend / divisor, nil
}

quotient := must.Return(divide(1, 2))
```
