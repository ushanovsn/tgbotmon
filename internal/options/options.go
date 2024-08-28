package options

import (
	"github.com/ushanovsn/golanglogger"
)

// Universal interface for use options
type Options interface {
	// Getting the logger interface object (the interface is actually a pointer)
	GetLogger() golanglogger.Golanglogger
	// Getting the logger logging level
	GetLoggerLevelParam() golanglogger.LoggingLevel
	// Set the logger object (an interface object or pointer to object that imlement interface Golanglogger)
	SetLogger(log golanglogger.Golanglogger)
	// Getting the loger file path (or just name)
	GetLogFileName() string

	// Getting the config file path (or just name)
	GetConfFileName() string
	// Getting the pointer to configurations structure
	GetConfigUniversalPtr() interface{}

	// Getting the description for config file
	GetConfigDescr() string
}
