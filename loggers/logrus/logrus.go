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

package logrus

import (
	"fmt"

	"github.com/Fishwaldo/go-logadapter"
	"github.com/sirupsen/logrus"
)

var _ logadapter.Logger = (*LruLogger)(nil)

type LruLogger struct {
	Lru *logrus.Entry
}

func (l LruLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.Lru.Debugf(msg, keysAndValues...)
}
func (l LruLogger) Error(msg string, keysAndValues ...interface{}) {
	l.Lru.Errorf(msg, keysAndValues...)
}
func (l LruLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.Lru.Fatalf(msg, keysAndValues...)
}
func (l LruLogger) Info(msg string, keysAndValues ...interface{}) {
	l.Lru.Infof(msg, keysAndValues...)
}
func (l LruLogger) Panic(msg string, keysAndValues ...interface{}) {
	l.Lru.Panicf(msg, keysAndValues...)
}
func (l LruLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.Lru.Warnf(msg, keysAndValues...)
}
func (l *LruLogger) With(key string, value interface{}) logadapter.Logger {
	nl := &LruLogger{Lru: logrus.NewEntry(l.Lru.Logger)}
	nl.Lru = l.Lru.WithField(key, value)
	return nl
}
func (l *LruLogger) Trace(key string, args ...interface{}) {
	l.Lru.Tracef(key, args...)
}
func (l *LruLogger) Sync() {
}
func (l *LruLogger) New(name string) logadapter.Logger {
	nl := &LruLogger{Lru: logrus.NewEntry(l.Lru.Logger)}
	nl.SetPrefix(name)
	return nl
}
func (l *LruLogger) SetPrefix(name string) {
	l.Lru = l.Lru.WithField("Prefix", name)
}

func (l *LruLogger) GetPrefix() (string) {
	prefix, ok := l.Lru.Data["Prefix"]
	if ok {
		return fmt.Sprint(prefix)
	}
	return ""
}
func (l *LruLogger) SetLevel(level logadapter.Log_Level) {
	switch level {
		case logadapter.LOG_TRACE:
			l.Lru.Logger.SetLevel(logrus.TraceLevel)
		case logadapter.LOG_DEBUG:
			l.Lru.Logger.SetLevel(logrus.DebugLevel)
		case logadapter.LOG_INFO:
			l.Lru.Logger.SetLevel(logrus.InfoLevel)
		case logadapter.LOG_WARN:
			l.Lru.Logger.SetLevel(logrus.WarnLevel)
		case logadapter.LOG_ERROR:
			l.Lru.Logger.SetLevel(logrus.ErrorLevel)
		case logadapter.LOG_FATAL:
			l.Lru.Logger.SetLevel(logrus.FatalLevel)
		case logadapter.LOG_PANIC:
			l.Lru.Logger.SetLevel(logrus.PanicLevel)		
	}
}
func (l *LruLogger) GetLevel() (logadapter.Log_Level) {
	switch l.Lru.Logger.GetLevel() {
	case logrus.TraceLevel:
		return logadapter.LOG_TRACE
	case logrus.DebugLevel:
		return logadapter.LOG_DEBUG
	case logrus.InfoLevel:
		return logadapter.LOG_INFO
	case logrus.WarnLevel:
		return logadapter.LOG_WARN
	case logrus.ErrorLevel:
		return logadapter.LOG_ERROR
	case logrus.FatalLevel:
		return logadapter.LOG_FATAL
	case logrus.PanicLevel:
		return logadapter.LOG_PANIC
	}
	return logadapter.LOG_UNKNOWN
}

//LogrusDefaultLogger Return Logger based on logrus with new instance
func LogrusDefaultLogger() (*LruLogger) {
	// TODO control verbosity
	l := &LruLogger{Lru: logrus.NewEntry(logrus.New())}
	return l
}

