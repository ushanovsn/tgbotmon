package options

import (
	"github.com/ushanovsn/golanglogger"
)

// Telegram Bot configuration values
//
// Tags "cfg" and "descr" uses for config file
type TgBotConfig struct {
	// host address
	SrvHost string `cfg:"server_host" descr:"Server host address"`
	// host port
	SrvPort uint `cfg:"server_port" descr:"Server host port"`
	// host port
	Token string `cfg:"tgbot_token" descr:"Token for Telegram Bot"`
	// logger level
	LogLevel string `cfg:"logging_level" descr:"Logger logging level (Debug/Info/Warning/Error)"`
	// log file size in megabytes
	LogSizeMb uint `cfg:"log_file_size_mb" descr:"Log file size in megabytes (0 - one file/no split)"`
	// log file size in megabytes
	LogSizeD uint `cfg:"log_file_size_day" descr:"Log file size in days (0 - one file/no split)"`
	// logging file (no file if empty string)
	LogFile string `cfg:"log_file" descr:"File name or full path for logging file (without spaces or use quotes)"`
	// file with configuration parameters
	ConfFile string `cfg:"config_file" descr:"Configuration file name or full path (without spaces or use quotes). When file not exist - it will be creating"`
}

// Telegram Bot object (full data of TgBot)
type TgBotObj struct {
	conf   TgBotConfig
	logger golanglogger.Golanglogger
}

// *********************   Interface "Options" implementation   *********************

// Getting the logger interface object (the interface is actually a pointer)
func (obj *TgBotObj) GetLogger() golanglogger.Golanglogger {
	return obj.logger
}

// Getting the logger logging level
func (obj *TgBotObj) GetLoggerLevelParam() golanglogger.LoggingLevel {
	v, _ := golanglogger.LoggingLevelValue(obj.conf.LogLevel)

	return v
}

// Set the logger object (an interface object or pointer to object that imlement interface Golanglogger)
func (obj *TgBotObj) SetLogger(log golanglogger.Golanglogger) {
	obj.logger = log
}

// Getting the loger file path (or just name)
func (obj *TgBotObj) GetLogFileName() string {
	return obj.conf.LogFile
}

// Getting the config file path (or just name)
func (obj *TgBotObj) GetConfFileName() string {
	return obj.conf.ConfFile
}

// Getting the universal pointer to configurations structure (as Interface)
func (obj *TgBotObj) GetConfigUniversalPtr() interface{} {
	return &obj.conf
}

// Getting the description for config file
func (obj *TgBotObj) GetConfigDescr() string {
	return DefSrvConfDescr
}

// *********************   Specific for tgbot methods   *********************

// Getting the pointer to TgBot configurations structure
func (obj *TgBotObj) GetConfigPtr() *TgBotConfig {
	return &obj.conf
}

// Getting the server port host address
func (obj *TgBotObj) GetServerHost() string {
	return obj.conf.SrvHost
}

// Getting the server port number
func (obj *TgBotObj) GetServerPort() uint {
	return obj.conf.SrvPort
}

// Getting the log file size mb
func (obj *TgBotObj) GetLogFSizeMb() uint {
	return obj.conf.LogSizeMb
}

// Getting the log file size days
func (obj *TgBotObj) GetLogFSizeD() uint {
	return obj.conf.LogSizeD
}

// Getting the token
func (obj *TgBotObj) GetToken() string {
	return obj.conf.Token
}
