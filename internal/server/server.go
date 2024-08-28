package server

import (
	//"fmt"
	"github.com/ushanovsn/tgbotmon/internal/options"

	"github.com/ushanovsn/golanglogger"
)



// Startimg servers processes
func StartServer(srv *options.ServerObj) {
	log := srv.GetLogger()
	log.Out("Server starting...")



	log.Out("Server successfully started!")
}


// Stop all process of server
func StopServer(srv *options.ServerObj) {
	// stopping logger
	srv.GetLogger().StopLog()
}



// Init server data and configurations.
//
// Load default values when no config found
func InitServer() *options.ServerObj {
	// create server object (options)
	var srv options.ServerObj
	// start config with default values
	srv.SetDefaultConf()
	// receive flags at start and use it
	setCmdFlags(srv.GetConfigPtr())
	
	// start logger with init values (flag received or default value)
	log := golanglogger.NewSync(srv.GetLoggerLevelParam(), srv.GetLogFileName())
	// save logger to server object
	srv.SetLogger(log)
	//log.SetName(options.DefSrvLogName)

	log.Out("The server is being initialized now...")

	// load and process configuration file
	ok := options.ProcConfig(&srv)
	if !ok {
		log.OutError("Error while read configuration from file. Missing parameters was set to default values")
	}

	setCmdFlags(srv.GetConfigPtr())
	log.OutInfo("Updated config by received flags")

	// apply the configuration
	log.Out("Now applying configuration parameters")

	if log.CurrentLevel() != srv.GetLoggerLevelParam() {
		log.SetLevel(srv.GetLoggerLevelParam())
	}

	return &srv
}
