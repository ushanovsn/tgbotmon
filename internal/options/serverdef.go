package options

const (
	// default config file name
	DefSrvConfFile string = "server.conf"
	// default log file name
	DefSrvLogFile string = "server.log"

	// default logger level
	DefSrvLogLvl string = "Error"
	// default logger name
	DefSrvLogName string = "SERVER"
	// default log file size in megabytes
	DefSrvLogSizeMb uint = 0
	// default log file size in days
	DefSrvLogSizeDay uint = 0

	// default host address
	DefSrvHost string = "localhost"
	// default host port
	DefSrvPort uint = 3003
	// default server GUI port
	DefSrvGuiPort uint = 3005

	// default GUI activating
	DefSrvUseGui bool = false

	// default description for config file
	DefSrvConfDescr string = "The configuration file for the \"SERVER\" process"
)

// Set default values to configuration structure
func (obj *ServerObj) SetDefaultConf() {
	obj.conf.Host = DefSrvHost
	obj.conf.Port = DefSrvPort
	obj.conf.UseGui = DefSrvUseGui
	obj.conf.GuiPort = DefSrvGuiPort
	obj.conf.LogLevel = DefSrvLogLvl
	obj.conf.LogSizeMb = DefSrvLogSizeMb
	obj.conf.LogSizeD = DefSrvLogSizeDay
	obj.conf.LogFile = DefSrvLogFile
	obj.conf.ConfFile = DefSrvConfFile
}
