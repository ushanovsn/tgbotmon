package main

import (
	"github.com/ushanovsn/tgbotmon/internal/server"
)



func main() {
	// load and init parameters
	conf, err := server.InitBot()

	if err == nil {
		// possible to start tg bot server
		server.StartBot(conf)
	}
}