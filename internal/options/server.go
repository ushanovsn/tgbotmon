package options

import (
	"github.com/ushanovsn/golanglogger"
)


// Server configuration values
type ServerConfig struct {
	// host address
	Host     string	`cfg:"host" descr:"Local host address"`
	// host port
	Port     uint	`cfg:"port" descr:"Local host port"`
	// gui enable flag
	UseGui   bool	`cfg:"gui_enable" descr:"Server GUI enable flag (true/false)"`
	// gui port (if enabled)
	GuiPort  uint	`cfg:"gui_port" descr:"Server GUI port"`
	// logger level
	LogLevel string	`cfg:"logging_level" descr:"Logger logging level (Debug/Info/Warning/Error)"`
	// logging file (no file if empty string)
	LogFile  string	`cfg:"log_file" descr:"File name or full path for logging file (without spaces or use quotes)"`
	// file with configuration parameters
	ConfFile string	`cfg:"config_file" descr:"Configuration file name or full path (without spaces or use quotes). When file not exist - it will be creating"`
}

// Server object (full data of server)
type ServerObj struct {
	conf   ServerConfig
	logger golanglogger.Golanglogger
}

// Getting the logger interface object (the interface is actually a pointer)
func (obj *ServerObj) GetLogger() golanglogger.Golanglogger {
	return obj.logger
}

// Getting the logger logging level
func (obj *ServerObj) GetLoggerLevel() golanglogger.LoggingLevel {
	return obj.logger.CurrentLevel()
}

// Set the logger object (an interface object or pointer to object that imlement interface Golanglogger)
func (obj *ServerObj) SetLogger(log golanglogger.Golanglogger) {
	obj.logger = log
}

// Getting the pointer to configurations structure
func (obj *ServerObj) GetConfigPtr() *ServerConfig {
	return &obj.conf
}

// Getting the loger file path (or just name)
func (obj *ServerObj) GetLogFile() string {
	return obj.conf.LogFile
}

// Getting the config file path (or just name)
func (obj *ServerObj) GetConfFile() string {
	return obj.conf.ConfFile
}
