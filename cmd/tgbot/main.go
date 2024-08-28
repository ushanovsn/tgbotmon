package main

import (
	"github.com/ushanovsn/tgbotmon/internal/tgbot"
	"time"
)

func main() {
	bot := tgbot.InitBot()
	tgbot.StartBot(bot)

	time.Sleep(3 * time.Second)

	tgbot.StopBot(bot)
}
