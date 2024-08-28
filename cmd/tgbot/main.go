package main

import (
	"time"
	"github.com/ushanovsn/tgbotmon/internal/tgbot"
)



func main() {
	bot := tgbot.InitBot()
	tgbot.StartBot(bot)

	time.Sleep(3 * time.Second)

	tgbot.StopBot(bot)
}