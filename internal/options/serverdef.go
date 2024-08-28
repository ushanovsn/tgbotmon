package options


const (
	// default config file name
	DefSrvConfFile string = "server.conf"
	// default log file name
	DefSrvLogFile string = "server.log"
	
	// default logger level
	//DefSrvLogLvl string = "Error"
	DefSrvLogLvl string = "Error"
	// default logger name
	DefSrvLogName string = "SERVER"

	// default host address
	DefSrvHost string = "localhost"
	// default host port
	DefSrvPort uint = 3003
	// default server GUI port
	DefSrvGuiPort uint = 3005

	// default GUI activating
	DefUseGui bool = false

	// default description for config file
	DefConfDescr string = "The configuration file for the \"SERVER\" process"
 )


 // Set default values to configuration structure
 func (obj *ServerObj) SetDefaultConf() {
	obj.conf.Host = DefSrvHost
	obj.conf.Port = DefSrvPort
	obj.conf.UseGui = DefUseGui
	obj.conf.GuiPort = DefSrvGuiPort
	obj.conf.LogLevel = DefSrvLogLvl
	obj.conf.LogFile = DefSrvLogFile
	obj.conf.ConfFile = DefSrvConfFile
}