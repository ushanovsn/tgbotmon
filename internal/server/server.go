package server

import (
	"fmt"

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
	log := golanglogger.NewSync(srv.GetLoggerLevel(), srv.GetLogFile())
	// save logger to server object
	srv.SetLogger(log)

	log.Out(fmt.Sprintf("Start to load configuration from \"%s\" file", srv.GetConfFile()))

	// load and process configuration file
	ok := getConfig(&srv)

	if !ok {
		log.OutError("Error while read configuration from file. Some parameters was set to default values")
	}

	// updating the flags values in to the loaded configuration - flags have a higher priority
	setCmdFlags(srv.GetConfigPtr())

	// apply the final configuration
	//.........................



	log.SetName(options.DefSrvLogFile)

	return &srv
}
