package main

import (
	"github.com/ushanovsn/tgbotmon/internal/tgbot"
)



func main() {
	// load and init parameters
	conf, err := tgbot.InitBot()

	if err == nil {
		// possible to start tg bot server
		tgbot.StartBot(conf)
	}
}