# strongoid

A simple package providing easy access to several function that cover the must-pattern, which is often utilized in Go.

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
```

The *must.Succeed(err error)* function internally calls *panic(err)* if a non-nil error is passed.
