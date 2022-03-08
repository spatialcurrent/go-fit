[![CircleCI](https://circleci.com/gh/spatialcurrent/go-fit/tree/main.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-fit/tree/main)
[![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-fit?style=flat-square)](https://goreportcard.com/report/github.com/spatialcurrent/go-fit)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/spatialcurrent/go-fit)](https://pkg.go.dev/github.com/spatialcurrent/go-fit)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-fit/blob/master/LICENSE)

# go-fit

# Description

**go-fit** is a library to fit data structures to their underlying types.

# Usage

**Go**

You can import **go-fit** as a library with:

```go
import (
  "github.com/spatialcurrent/go-fit/pkg/fit"
)
```

The easiest pattern is to use the `fit.Fit(in interface{}) interface{}` function.

```go
fittedObject := fit.Fit(obj)
```

See [fit](https://pkg.go.dev/github.com/spatialcurrent/go-fit/pkg/fit) in the docs for information on how to use Go API.

# Testing

To run Go tests using `make test_go` or (`bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-fit/blob/main/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
