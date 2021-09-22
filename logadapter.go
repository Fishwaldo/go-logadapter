package logadapter

import (
)

type Logger interface {
	Trace(message string, params ...interface{})
	Debug(message string, params ...interface{})
	Info(message string, params ...interface{})
	Warn(message string, params ...interface{})
	Error(message string, params ...interface{})
	Fatal(message string, params ...interface{})
	Panic(message string, params ...interface{})
	New(name string) (l Logger)
	With(key string, value interface{}) (l Logger)
	SetPrefix(name string)
	GetPrefix() (string)
	Sync()
}


