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

package logrus_test

import (
	// "bytes"
	// "os"
	// "regexp"
	"testing"

//	"github.com/Fishwaldo/go-logadapter/loggers/zap"
)

func TestDefaultLogger(t *testing.T) {
	//_ := zaplog.ZapLogger{}
	//logger.SetLevel(stdlogger.LOG_TRACE)
	//if logger.GetLevel() != stdlogger.LOG_TRACE {
	//	t.Errorf("Can't Set Logging Level")
	//}
}

// func captureOutput(l *zaplog.ZapLogger, f func()) string {
//     // var buf bytes.Buffer
//     // l.Log.SetOutput(&buf)
//     // f()
//     // l.Log.SetOutput(os.Stderr)
//     // return buf.String()
// }
func TestLogTrace(t *testing.T) {
	// logger := stdlogger.DefaultLogger()
	// logger.SetLevel(stdlogger.LOG_TRACE)
	// output := captureOutput(logger, func() { 
	// 	logger.Trace("Hello %s", "world")
	// })
	// validmsg := regexp.MustCompile(`^.* TRACE: Hello world \- {}`)
	// if !validmsg.MatchString(output) {
	// 	t.Errorf("Log Trace Failed: %s", output)
	// }
}
func TestLogDebug(t *testing.T) {
	// logger := stdlogger.DefaultLogger()
	// logger.SetLevel(stdlogger.LOG_TRACE)
	// output := captureOutput(logger, func() { 
	// 	logger.Debug("Hello %s", "world")
	// })
	// validmsg := regexp.MustCompile(`^.* DEBUG: Hello world \- {}`)
	// if !validmsg.MatchString(output) {
	// 	t.Errorf("Log Debug Failed: %s", output)
	// }
}

func TestLogInfo(t *testing.T) {
	// logger := stdlogger.DefaultLogger()
	// logger.SetLevel(stdlogger.LOG_TRACE)
	// output := captureOutput(logger, func() { 
	// 	logger.Info("Hello %s", "world")
	// })
	// validmsg := regexp.MustCompile(`^.* INFO: Hello world \- {}`)
	// if !validmsg.MatchString(output) {
	// 	t.Errorf("Log Info Failed: %s", output)
	// }
}

func TestLogWarn(t *testing.T) {
	// logger := stdlogger.DefaultLogger()
	// logger.SetLevel(stdlogger.LOG_TRACE)
	// output := captureOutput(logger, func() { 
	// 	logger.Warn("Hello %s", "world")
	// })
	// validmsg := regexp.MustCompile(`^.* WARN: Hello world \- {}`)
	// if !validmsg.MatchString(output) {
	// 	t.Errorf("Log Warn Failed: %s", output)
	// }
}

func TestLogError(t *testing.T) {
	// logger := stdlogger.DefaultLogger()
	// logger.SetLevel(stdlogger.LOG_TRACE)
	// output := captureOutput(logger, func() { 
	// 	logger.Error("Hello %s", "world")
	// })
	// validmsg := regexp.MustCompile(`^.* ERROR: Hello world \- {}`)
	// if !validmsg.MatchString(output) {
	// 	t.Errorf("Log Error Failed: %s", output)
	// }
}

func TestLogFatal(t *testing.T) {
	//logger := DefaultLogger()
	//logger.SetLevel(LOG_TRACE)
	//output := captureOutput(logger, func() { 
	//	logger.Fatal("Hello %s", "world")
	//})
	//validmsg := regexp.MustCompile(`^.* FATAL: Hello world \- {}`)
	//if !validmsg.MatchString(output) {
	//	t.Errorf("Log Fatal Failed: %s", output)
	//}
}

func TestLogPanic(t *testing.T) {
	// logger := stdlogger.DefaultLogger()
	// logger.SetLevel(stdlogger.LOG_TRACE)
	// defer func() {
	// 	if err := recover(); err == nil {
	// 		t.Errorf("Log Panic Recovery Failed")
	// 	}
	// }()
	// output := captureOutput(logger, func() { 
	// 	logger.Panic("Hello %s", "world")
	// })

	// validmsg := regexp.MustCompile(`^.* PANIC: Hello world \- {}`)
	// if !validmsg.MatchString(output) {
	// 	t.Errorf("Log Panic Failed: %s", output)
	// }
}
