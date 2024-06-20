package server

import (
	"flag"
)

// structure for holding flags values
type flagStruct struct {
	logLevel string
	logFile string
	confFile string
}

// processing flags values
func getFlags() flagStruct{
	var flags flagStruct

	_ = flag.Value(&o.Net)

	flag.Var(&o.Net, "a", "Server net address host:port")
	flag.StringVar(&o.Logger.Level, "l", "info", "Logging level")

	flag.Parse()

	return flags
}
