package options

import (
	"github.com/ushanovsn/golanglogger"
)


// Universal interface for use options
type Options interface {
	// Getting the logger interface object (the interface is actually a pointer)
	GetLogger() golanglogger.Golanglogger
	
}