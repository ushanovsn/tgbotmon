package server

import (
	"flag"
	"github.com/ushanovsn/tgbotmon/internal/options"
	"os"
)

// Set (update) flags values from cmd into config.
//
// Update only received flags, others configuration parameters are not changes.
func setCmdFlags(conf *options.ServerConfig) {
	var tmpF options.ServerConfig
	var flags *flag.FlagSet

	// if base flags was set and parsed - changing flagset
	if !flag.Parsed() {
		flags = flag.CommandLine
	} else {
		flags = flag.NewFlagSet("new flag set", flag.ExitOnError)
	}

	// deffine the parameters
	flags.StringVar(&tmpF.Host, "h", conf.Host, "Server host address")
	flags.UintVar(&tmpF.Port, "p", conf.Port, "Server port")
	flags.BoolVar(&tmpF.UseGui, "gui", conf.UseGui, "Enable server GUI (bool)")
	flags.UintVar(&tmpF.GuiPort, "guiport", conf.GuiPort, "Server GUI port")
	flags.StringVar(&tmpF.LogLevel, "loglvl", conf.LogLevel, "Server logging level (text)")
	flags.StringVar(&tmpF.LogFile, "logfile", conf.LogFile, "Server log file path\\name")
	flags.StringVar(&tmpF.ConfFile, "conffile", conf.ConfFile, "Server config file path\\name")

	// service will stopping in this place when error occurs
	_ = flags.Parse(os.Args[1:])

	// now set parameters to config
	conf.Host = tmpF.Host
	conf.Port = tmpF.Port
	conf.UseGui = tmpF.UseGui
	conf.GuiPort = tmpF.GuiPort
	conf.LogLevel = tmpF.LogLevel
	conf.LogFile = tmpF.LogFile
	conf.ConfFile = tmpF.ConfFile
}
