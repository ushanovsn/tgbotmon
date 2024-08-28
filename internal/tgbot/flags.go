package tgbot

import (
	"flag"
	"os"
	"github.com/ushanovsn/tgbotmon/internal/options"
)


// Set (update) flags values from cmd into tgbot config.
//
// Update only received flags, others configuration parameters are not changes.
func setCmdFlags(conf *options.TgBotConfig){
	var tmpF options.TgBotConfig
	var flags *flag.FlagSet

	// if base flags was set and parsed - changing flagset
	if !flag.Parsed() {
		flags = flag.CommandLine
	} else {
		flags = flag.NewFlagSet("new flag set", flag.ExitOnError)
	}

	// deffine the parameters
	flags.StringVar(&tmpF.SrvHost, "sh", conf.SrvHost, "Server host address")
	flags.UintVar(&tmpF.SrvPort, "sp", conf.SrvPort, "Server port")
	flags.StringVar(&tmpF.Token, "t", conf.Token, "Telegram Bot token")
	flags.StringVar(&tmpF.LogLevel, "loglvl", conf.LogLevel, "Server logging level (text)")
	flags.StringVar(&tmpF.LogFile, "logfile", conf.LogFile, "Server log file path\\name")
	flags.StringVar(&tmpF.ConfFile, "conffile", conf.ConfFile, "Server config file path\\name")

	// service will stopping in this place when error occurs
	flags.Parse(os.Args[1:])

	// now set parameters to config
	conf.SrvHost = tmpF.SrvHost
	conf.SrvPort = tmpF.SrvPort
	conf.Token = tmpF.Token
	conf.LogLevel = tmpF.LogLevel
	conf.LogFile = tmpF.LogFile
	conf.ConfFile = tmpF.ConfFile
}
