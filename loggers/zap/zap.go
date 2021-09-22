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

package zaplog

import (
	"github.com/Fishwaldo/go-logadapter"
	"go.uber.org/zap"
)

var _ logadapter.Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	Zap *zap.SugaredLogger
	lvl logadapter.Log_Level
}

func (l *ZapLogger) With(key string, value interface{}) logadapter.Logger {
	nl := &ZapLogger{Zap: l.Zap.With(key, value)}
	return nl
}

func (l *ZapLogger) Trace(message string, params ...interface{}) {
	if l.lvl <= logadapter.LOG_TRACE {
	l.Zap.Debugf(message, params...)
	}
}
func (l *ZapLogger) Debug(message string, params ...interface{}) {
	if l.lvl <= logadapter.LOG_DEBUG {
		l.Zap.Debugf(message, params...)
	}
}
func (l *ZapLogger) Info(message string, params ...interface{}) {
	if l.lvl <= logadapter.LOG_INFO {
	l.Zap.Infof(message, params...)
	}
}
func (l *ZapLogger) Warn(message string, params ...interface{}) {
	if l.lvl <= logadapter.LOG_WARN {
		l.Zap.Warnf(message, params...)
	}
}
func (l *ZapLogger) Error(message string, params ...interface{}) {
	if l.lvl <= logadapter.LOG_ERROR {
		l.Zap.Errorf(message, params...)
	}
}
func (l *ZapLogger) Fatal(message string, params ...interface{}) {
	if l.lvl <= logadapter.LOG_FATAL {
		l.Zap.Fatalf(message, params...)
	}
}
func (l *ZapLogger) Panic(message string, params ...interface{}) {
	if l.lvl <= logadapter.LOG_PANIC {
		l.Zap.Panicf(message, params...)
	}
}
func (l *ZapLogger) New(name string) (nl logadapter.Logger) {
	zl := &ZapLogger{Zap: l.Zap}
	zl.Zap.Named(name)
	return zl
}

func (l *ZapLogger) Sync() {

}

func (l *ZapLogger) SetPrefix(name string) {
	// XXX TODO
}
func (l ZapLogger) GetPrefix() (string) {
	return ""
}

func (l *ZapLogger) GetLevel() (logadapter.Log_Level) {
	return l.lvl
}
func (l *ZapLogger) SetLevel(lvl logadapter.Log_Level) {
	l.lvl = lvl
}

//DefaultLogger Return Default Sched Logger based on Zap's sugared logger
func NewZapLogger() *ZapLogger {
	// TODO control verbosity
	loggerBase, _ := zap.NewDevelopment(zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	sugarLogger := loggerBase.Sugar()
	return &ZapLogger{
		Zap: sugarLogger,
	}
}
