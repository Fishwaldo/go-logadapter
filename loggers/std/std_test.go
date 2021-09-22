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

package stdlogger_test

import (
	"bytes"
	"os"
	"regexp"
	"testing"

	"github.com/Fishwaldo/go-logadapter"
	"github.com/Fishwaldo/go-logadapter/loggers/std"
)

func TestStdLogger(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	if logger.GetLevel() != logadapter.LOG_TRACE {
		t.Errorf("Can't Set Logging Level")
	}
}

func captureOutput(l *stdlogger.StdLogger, f func()) string {
    var buf bytes.Buffer
    l.Log.SetOutput(&buf)
    f()
    l.Log.SetOutput(os.Stderr)
    return buf.String()
}
func TestStdTrace(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	output := captureOutput(logger, func() { 
		logger.Trace("Hello %s", "world")
	})
	validmsg := regexp.MustCompile(`^.* TRACE: Hello world \- {}`)
	if !validmsg.MatchString(output) {
		t.Errorf("Log Trace Failed: %s", output)
	}
}
func TestStdDebug(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	output := captureOutput(logger, func() { 
		logger.Debug("Hello %s", "world")
	})
	validmsg := regexp.MustCompile(`^.* DEBUG: Hello world \- {}`)
	if !validmsg.MatchString(output) {
		t.Errorf("Log Debug Failed: %s", output)
	}
}

func TestStdInfo(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	output := captureOutput(logger, func() { 
		logger.Info("Hello %s", "world")
	})
	validmsg := regexp.MustCompile(`^.* INFO: Hello world \- {}`)
	if !validmsg.MatchString(output) {
		t.Errorf("Log Info Failed: %s", output)
	}
}

func TestStdWarn(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	output := captureOutput(logger, func() { 
		logger.Warn("Hello %s", "world")
	})
	validmsg := regexp.MustCompile(`^.* WARN: Hello world \- {}`)
	if !validmsg.MatchString(output) {
		t.Errorf("Log Warn Failed: %s", output)
	}
}

func TestStdError(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	output := captureOutput(logger, func() { 
		logger.Error("Hello %s", "world")
	})
	validmsg := regexp.MustCompile(`^.* ERROR: Hello world \- {}`)
	if !validmsg.MatchString(output) {
		t.Errorf("Log Error Failed: %s", output)
	}
}

func TestStdFatal(t *testing.T) {
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

func TestStdPanic(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Log Panic Recovery Failed")
		}
	}()
	output := captureOutput(logger, func() { 
		logger.Panic("Hello %s", "world")
	})

	validmsg := regexp.MustCompile(`^.* PANIC: Hello world \- {}`)
	if !validmsg.MatchString(output) {
		t.Errorf("Log Panic Failed: %s", output)
	}
}
func TestStdPrefix(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	if logger.GetPrefix() != "log" {
		t.Errorf("Empty Log Prefix Failed: %s", logger.GetPrefix())
	}
	logger.SetPrefix("TestStd")

	output := captureOutput(logger, func() { 
		logger.Info("Hello %s", "world")
	})
	validmsg := regexp.MustCompile(`^TestStd .* INFO: Hello world \- {}`)
	if !validmsg.MatchString(output) {
		t.Errorf("Log Prefix Failed: %s", output)
	}
	if logger.GetPrefix() != "TestStd" {
		t.Errorf("Log Prefix Failed: %s", logger.GetPrefix())
	}
}

func TestStdLevel(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	if logger.GetLevel() != logadapter.LOG_TRACE {
		t.Errorf("Log SetLevel to Trace Failed: %d", logger.GetLevel())
	}
	logger.SetLevel(logadapter.LOG_DEBUG)
	if logger.GetLevel() != logadapter.LOG_DEBUG {
		t.Errorf("Log SetLevel to Debug Failed: %d", logger.GetLevel())
	}
	logger.SetLevel(logadapter.LOG_INFO)
	if logger.GetLevel() != logadapter.LOG_INFO {
		t.Errorf("Log SetLevel to Info Failed: %d", logger.GetLevel())
	}
	logger.SetLevel(logadapter.LOG_WARN)
	if logger.GetLevel() != logadapter.LOG_WARN {
		t.Errorf("Log SetLevel to Trace Failed: %d", logger.GetLevel())
	}
	logger.SetLevel(logadapter.LOG_ERROR)
	if logger.GetLevel() != logadapter.LOG_ERROR {
		t.Errorf("Log SetLevel to Error Failed: %d", logger.GetLevel())
	}
	logger.SetLevel(logadapter.LOG_FATAL)
	if logger.GetLevel() != logadapter.LOG_FATAL {
		t.Errorf("Log SetLevel to Fatal Failed: %d", logger.GetLevel())
	}
	logger.SetLevel(logadapter.LOG_PANIC)
	if logger.GetLevel() != logadapter.LOG_PANIC {
		t.Errorf("Log SetLevel to Panic Failed: %d", logger.GetLevel())
	}
}

func TestStdNewLog(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	logger2 := logger.New("Blah")
	output := captureOutput(logger2.(*stdlogger.StdLogger), func() { 
		logger2.Error("Hello %s", "world")
	})
	validmsg := regexp.MustCompile(`^Blah .* ERROR: Hello world \- {}`)
	if !validmsg.MatchString(output) {
		t.Errorf("Log Error Failed: %s", output)
	}
}

func TestStdWith(t *testing.T) {
	logger := stdlogger.DefaultLogger()
	logger.SetLevel(logadapter.LOG_TRACE)
	logger2 := logger.With("Test", "Message")

	output := captureOutput(logger2.(*stdlogger.StdLogger), func() { 
		logger2.Info("Hello %s", "world")
	})
	validmsg := regexp.MustCompile(`^.* INFO: Hello world \- {\"Test\":\"Message\"}`)
	if !validmsg.MatchString(output) {
		t.Errorf("Log With Failed: %s", output)
	}
}