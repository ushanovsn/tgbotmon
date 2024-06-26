package server

import (
	"flag"
	"github.com/ushanovsn/tgbotmon/internal/options"
)


// Set (update) flags values from cmd into server config.
// Update only received flags, others configuration parameters are not changes.
func setCmdFlags(conf *options.ServerConfig){

	flag.StringVar(&conf.Host, "host", conf.Host, "Server host address")
	flag.UintVar(&conf.Port, "port", conf.Port, "Server port")
	flag.BoolVar(&conf.UseGui, "gui", conf.UseGui, "Enable server GUI")
	flag.UintVar(&conf.GuiPort, "guiport", conf.GuiPort, "Server GUI port")
	flag.StringVar(&conf.LogLevel, "loglvl", conf.LogLevel, "Server logging level")
	flag.StringVar(&conf.LogFile, "logfilename", conf.LogFile, "Server log file name")
	flag.StringVar(&conf.ConfFile, "confFile", conf.ConfFile, "Server config file path")
}
