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

package logadapter_test

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

func Example() {
	ts := TestStruct{}
	ts.Init()
	ts.LogSomething()
	ts.StructuredLog()
	ts.Logger.Warn("This is outside the Structure")
}