# consterr

[![Build Status](https://travis-ci.com/Devoter/consterr.svg?branch=master)](https://travis-ci.com/Devoter/consterr)

This package provides a trivial implementation of Go `Error` interface, which cannot be modified by other code.

# Usage example

```go
import "github.com/Devoter/consterr"

const ErrSomethingWentWrong = consterr.Error("Something went wrong")
```
