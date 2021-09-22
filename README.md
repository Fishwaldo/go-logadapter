# logadapter

[![codecov](https://codecov.io/gh/Fishwaldo/go-logadapter/branch/master/graph/badge.svg)](https://codecov.io/gh/Fishwaldo/go-logadapter)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/Fishwaldo/go-logadapter)

## Sub Packages

* [utils](./utils)

## Examples

```golang
package main

import (
	"github.com/Fishwaldo/go-logadapter"
)

type TestStruct struct {
	Logger logadapter.Logger
}

func (t *TestStruct) Init() {
	temp := logadapter.DefaultLogger()
	temp.SetLevel(logadapter.LOG_TRACE)
	t.Logger = temp
}

func (t *TestStruct) LogSomething() {
	t.Logger.Info("This is a message")
}

func (t *TestStruct) StructuredLog() {
	t.Logger.With("Test", "Message").Trace("Hello %s", string("world"))
}

func main() {
	ts := TestStruct{}
	ts.Init()
	ts.LogSomething()
	ts.StructuredLog()
	ts.Logger.Warn("This is outside the Structure")
}

```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
