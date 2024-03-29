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

package stdlogger

import (
	"log"
	"os"
	"sync"
	"fmt"
	"strings"
	"encoding/json"

	"github.com/Fishwaldo/go-logadapter"
	"github.com/Fishwaldo/go-logadapter/internal/utils"
)

var _ logadapter.Logger = (*StdLogger)(nil)

//DefaultLogger uses Golang Standard Logging Libary
func DefaultLogger() (l *StdLogger) {
	stdlogger := log.New(os.Stderr, "log  - ", log.LstdFlags)
	stdlog := &StdLogger{Log: *stdlogger, keys: make(map[string]interface{})}
	return stdlog
}

type StdLogger struct {
	Log   log.Logger
	keys  map[string]interface{}
	mx    sync.Mutex
	level logadapter.Log_Level
}

func (l *StdLogger) Trace(message string, params ...interface{}) {
	if l.level <= logadapter.LOG_TRACE {
		l.Log.Printf("TRACE: %s - %s", fmt.Sprintf(message, params...), l.getKeys())
	}
}
func (l *StdLogger) Debug(message string, params ...interface{}) {
	if l.level <= logadapter.LOG_DEBUG {
		l.Log.Printf("DEBUG: %s - %s", fmt.Sprintf(message, params...), l.getKeys())
	}
}
func (l *StdLogger) Info(message string, params ...interface{}) {
	if l.level <= logadapter.LOG_INFO {
		l.Log.Printf("INFO: %s - %s", fmt.Sprintf(message, params...), l.getKeys())
	}
}
func (l *StdLogger) Warn(message string, params ...interface{}) {
	if l.level <= logadapter.LOG_WARN {
		l.Log.Printf("WARN: %s - %s", fmt.Sprintf(message, params...), l.getKeys())
	}
}
func (l *StdLogger) Error(message string, params ...interface{}) {
	if l.level <= logadapter.LOG_ERROR {
		l.Log.Printf("ERROR: %s - %s", fmt.Sprintf(message, params...), l.getKeys())
	}
}
func (l *StdLogger) Fatal(message string, params ...interface{}) {
  	if l.level <= logadapter.LOG_FATAL {
	l.Log.Fatal(fmt.Printf("FATAL: %s - %s", fmt.Sprintf(message, params...), l.getKeys()))
    }
}
func (l *StdLogger) Panic(message string, params ...interface{}) {
  	if l.level <= logadapter.LOG_PANIC {
	l.Log.Panic(fmt.Printf("PANIC: %s - %s", fmt.Sprintf(message, params...), l.getKeys()))
    }
}
func (l *StdLogger) New(name string) logadapter.Logger {
	//nl := &StdLogger{keys: l.keys}
	nl := &StdLogger{level: l.level, keys: make(map[string]interface{})}
	nl.SetPrefix(name)
	nl.Log.SetFlags(l.Log.Flags())
	nl.Log.SetOutput(l.Log.Writer())
	return nl
}
func (l *StdLogger) With(key string, value interface{}) logadapter.Logger {
	l.mx.Lock()
	defer l.mx.Unlock()
	stdlog := &StdLogger{level: l.level, keys: utils.CopyableMap(l.keys).DeepCopy()}
	stdlog.SetPrefix(l.GetPrefix())
	stdlog.Log.SetFlags(l.Log.Flags())
	stdlog.Log.SetOutput(l.Log.Writer())
	stdlog.keys[key] = value
	return stdlog
}
func (l *StdLogger) Sync() {
	// do nothing
}

func (l *StdLogger) getKeys() (message string) {
	l.mx.Lock()
	defer l.mx.Unlock()
	msg, err := json.Marshal(l.keys)
	if err == nil {
		return string(msg)
	}
	return err.Error()
}

func (l *StdLogger) SetLevel(level logadapter.Log_Level) {
	l.level = level
}

func (l *StdLogger) GetLevel() (logadapter.Log_Level) {
	return l.level
}

func (l *StdLogger) SetPrefix(name string) {
	l.Log.SetPrefix(fmt.Sprintf("%s - ", name))
}

func (l *StdLogger) GetPrefix() (string) {
	temp := strings.Fields(l.Log.Prefix())
	if len(temp) > 0 {
		return temp[0]
	} 
	return ""
}