package options

import (
	"github.com/ushanovsn/golanglogger"
)


// universal interface for interactions
type Options interface {
	GetLogger() golanglogger.Golanglogger
	
}