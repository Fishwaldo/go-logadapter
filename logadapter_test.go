package logadapter_test

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

func Example() {
	ts := TestStruct{}
	ts.Init()
	ts.LogSomething()
	ts.StructuredLog()
	ts.Logger.Warn("This is outside the Structure")
}