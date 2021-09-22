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

package logadapter

import (
)
// Logger is a interface that Applications/Libraries use
// 
// The actual Loggers should implement a structure 
// that Adhears to this Interface
type Logger interface {
	// Log a Trace Message
	Trace(message string, params ...interface{})
	// Log a Debug Message
	Debug(message string, params ...interface{})
	// Log a Info Message
	Info(message string, params ...interface{})
	// Log a Warn Message
	Warn(message string, params ...interface{})
	// Log a Error Message
	Error(message string, params ...interface{})
	// Log a Fatal Message (some implementations may call os.exit() here)
	Fatal(message string, params ...interface{})
	// Log a Panic Message (some implmentations may call Panic)
	Panic(message string, params ...interface{})
	// Create a New Logger Instance with Name
	New(name string) (l Logger)
	// Add Key/Value Pairs for Structured Logging and return a new Logger
	With(key string, value interface{}) (l Logger)
	// Set the Log Prefix
	SetPrefix(name string)
	// Get the Log Prefix
	GetPrefix() (string)
	// Set Logging Level
	SetLevel(Log_Level)
	// Get Logging Level
	GetLevel() (Log_Level)
	// Sync/Flush the Log Buffers 
	Sync()
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
	LOG_UNKNOWN
)