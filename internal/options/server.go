package options

import (
	"github.com/ushanovsn/golanglogger"
)


// server configuration
type ServerConfig struct {
	Host     string
	Port     uint
	UseGui   bool
	GuiPort  uint
	LogLevel string
	LogFile  string
	ConfFile string
}

// full server data
type ServerObj struct {
	conf   ServerConfig
	logger golanglogger.Golanglogger
}


func (obj *ServerObj) GetLogger() golanglogger.Golanglogger {
	return obj.logger
}

func (obj *ServerObj) GetLoggerLevel() golanglogger.LoggingLevel {
	return obj.logger.CurrentLevel()
}

func (obj *ServerObj) SetLogger(log golanglogger.Golanglogger) {
	obj.logger = log
}


func (obj *ServerObj) GetConfigPtr() *ServerConfig {
	return &obj.conf
}

func (obj *ServerObj) GetLogFile() string {
	return obj.conf.LogFile
}

func (obj *ServerObj) GetConfFile() string {
	return obj.conf.ConfFile
}
