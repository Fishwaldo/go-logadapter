/*
Package go-logadapter allows you to use different Log Libraries in your
packages/applications

This package just provides wrappers around different log libraries and 
allows the user of your library to use their preferred Logging Library

It implements common Logging Levels:

	* Trace
	* Debug
	* Info
	* Warn
	* Error
	* Fatal
	* Panic

It also supports basic structured logging features, but specifing a key/value 
pair using the With Statement.

Supported Logging Libraries

logadapter currently supports:

	* (std library logger) https://pkg.go.dev/log
	* (logrus) https://github.com/sirupsen/logrus
	* (zap) https://github.com/uber-go/zap

Adding additional Logging Libraries is realatively straight forward by creating
a wrapper that implements the (Logger) https://pkg.go.dev/github.com/Fishwaldo/go-logadapter#Logger Interface

*/
package logadapter