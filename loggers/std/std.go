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
	level Log_Level
}

type Log_Level int

const (
	LOG_TRACE Log_Level = iota
	LOG_DEBUG
	LOG_INFO
	LOG_WARN
	LOG_ERROR
	LOG_FATAL
	LOG_PANIC
)

func (l *StdLogger) Trace(message string, params ...interface{}) {
	if l.level <= LOG_TRACE {
		l.Log.Printf("TRACE: %s - %s", fmt.Sprintf(message, params...), l.getKeys())
	}
}
func (l *StdLogger) Debug(message string, params ...interface{}) {
	if l.level <= LOG_DEBUG {
		l.Log.Printf("DEBUG: %s - %s", fmt.Sprintf(message, params...), l.getKeys())
	}
}
func (l *StdLogger) Info(message string, params ...interface{}) {
	if l.level <= LOG_INFO {
		l.Log.Printf("INFO: %s - %s", fmt.Sprintf(message, params...), l.getKeys())
	}
}
func (l *StdLogger) Warn(message string, params ...interface{}) {
	if l.level <= LOG_WARN {
		l.Log.Printf("WARN: %s - %s", fmt.Sprintf(message, params...), l.getKeys())
	}
}
func (l *StdLogger) Error(message string, params ...interface{}) {
	if l.level <= LOG_ERROR {
		l.Log.Printf("ERROR: %s - %s", fmt.Sprintf(message, params...), l.getKeys())
	}
}
func (l *StdLogger) Fatal(message string, params ...interface{}) {
	l.Log.Fatal(fmt.Printf("FATAL: %s - %s", fmt.Sprintf(message, params...), l.getKeys()))
}
func (l *StdLogger) Panic(message string, params ...interface{}) {
	l.Log.Panic(fmt.Printf("PANIC: %s - %s", fmt.Sprintf(message, params...), l.getKeys()))
}
func (l *StdLogger) New(name string) logadapter.Logger {
	//nl := &StdLogger{keys: l.keys}
	nl := &StdLogger{level: l.level}
	nl.Log.SetPrefix(fmt.Sprintf("%s.%s", l.Log.Prefix(), name))
	nl.Log.SetFlags(l.Log.Flags())
	nl.Log.SetOutput(l.Log.Writer())
	return nl
}
func (l *StdLogger) With(key string, value interface{}) logadapter.Logger {
	l.mx.Lock()
	defer l.mx.Unlock()
	stdlog := &StdLogger{level: l.level, keys: utils.CopyableMap(l.keys).DeepCopy()}
	stdlog.Log.SetPrefix(l.Log.Prefix())
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

func (l *StdLogger) SetLevel(level Log_Level) {
	l.level = level
}

func (l *StdLogger) GetLevel() (level Log_Level) {
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