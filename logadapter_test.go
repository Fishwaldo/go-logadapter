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