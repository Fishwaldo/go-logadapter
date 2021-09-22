# logadapter

[![codecov](https://codecov.io/gh/Fishwaldo/go-logadapter/branch/master/graph/badge.svg)](https://codecov.io/gh/Fishwaldo/go-logadapter)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/Fishwaldo/go-logadapter)

Package go-logadapter allows you to use different Log Libraries in your
packages/applications

This package just provides wrappers around different log libraries and
allows the user of your library to use their preferred Logging Library

It implements common Logging Levels:

```go
* Trace
* Debug
* Info
* Warn
* Error
* Fatal
* Panic
```

It also supports basic structured logging features, but specifing a key/value
pair using the With Statement.

## Supported Logging Libraries

logadapter currently supports:

```go
* [std library logger](https://pkg.go.dev/log)
* [logrus](https://github.com/sirupsen/logrus)
* [zap](https://github.com/uber-go/zap)
```

Adding additional Logging Libraries is realatively straight forward by creating
a wrapper that implements the [Logger](https://pkg.go.dev/github.com/Fishwaldo/go-logadapter#Logger) Interface

## Examples

```golang
/*
MIT License

Copyright (c) 2021 Justin Hammond

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"github.com/Fishwaldo/go-logadapter"
	"github.com/Fishwaldo/go-logadapter/loggers/std"
)

type TestStruct struct {
	Logger logadapter.Logger
}

func (t *TestStruct) Init() {
	temp := stdlogger.DefaultLogger()
	temp.SetLevel(stdlogger.LOG_TRACE)
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
