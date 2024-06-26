package server

import (
	"fmt"

	"github.com/ushanovsn/tgbotmon/internal/options"

	"github.com/ushanovsn/golanglogger"
)



func StartServer() {

	srv := initServer()

	log := srv.GetLogger()

	log.Out("Server starting...")

}

// Init server data and configurations.
// Load default values when no config found
func initServer() *options.ServerObj {
	// create full server data
	var srvOpt options.ServerObj

	// config with default values
	srvOpt.SetDefaultConf()

	// receive flags at start and update config
	setCmdFlags(srvOpt.GetConfigPtr())

	// start logger with init values (flag received or default value)
	log := golanglogger.NewSync(srvOpt.GetLoggerLevel(), srvOpt.GetLogFile())
	log.SetName(options.DefSrvLogFile)
	srvOpt.SetLogger(log)

	log.Out(fmt.Sprintf("Start to load configuration from \"%s\" file", srvOpt.GetConfFile()))

	// load and process configuration file
	ok := configurationProcess(&srvOpt)

	if !ok {
		log.OutError("Error while read configuration from file. Some parameters was set to default values")
	}

	// update flags values to loaded configuration
	setCmdFlags(srvOpt.GetConfigPtr())

	return &srvOpt
}
