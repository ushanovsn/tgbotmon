package tgbot

import (
	//"fmt"
	"fmt"

	"github.com/ushanovsn/tgbotmon/internal/options"

	"github.com/ushanovsn/golanglogger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Startimg servers processes
func StartBot(bot *options.TgBotObj) {
	log := bot.GetLogger()
	log.Out("TgBot starting...")

	InitializingBot(bot)

	log.Out("TgBot successfully started!")
}

// Stop all process of TgBot
func StopBot(bot *options.TgBotObj) {
	// stopping logger
	bot.GetLogger().StopLog()
}

// Init tgbot data and configurations.
//
// Load default values when no config found
func InitBot() *options.TgBotObj {
	// create bot object (options)
	var bot options.TgBotObj
	// start config with default values
	bot.SetDefaultConf()
	// receive flags at start and use it
	setCmdFlags(bot.GetConfigPtr())

	// start logger with init values (flag received or default value)
	log := golanglogger.NewSync(bot.GetLoggerLevelParam(), bot.GetLogFileName())
	// save logger to tgbot object
	bot.SetLogger(log)
	//log.SetName(options.DefBotLogName)

	log.Out("TgBot is being initialized now...")

	// load and process configuration file
	ok := options.ProcConfig(&bot)
	if !ok {
		log.OutError("Error while read configuration from file. Missing parameters was set to default values")
	}

	setCmdFlags(bot.GetConfigPtr())
	log.OutInfo("Updated config by received flags")

	// apply the configuration
	log.Out("Now applying configuration parameters")

	if log.CurrentLevel() != bot.GetLoggerLevelParam() {
		log.SetLevel(bot.GetLoggerLevelParam())
	}
	if szm, szd := log.CurrentFileControl(); szm != int(bot.GetLogFSizeMb()) || szd != int(bot.GetLogFSizeD()) {
		log.SetFileParam(int(szm), int(szd))
	}

	return &bot
}

func InitializingBot(bot *options.TgBotObj) {
	log := bot.GetLogger()

	tgBot, err := tgbotapi.NewBotAPI(bot.GetToken())
	if err != nil {
		log.OutError(fmt.Sprintf("Error when API init: %s", err.Error()))
	}

	log.Out(fmt.Sprintf("Bot authorized on account: %s", tgBot.Self.UserName))

	// updConf - структура с конфигом для получения апдейтов (0 - информируем телеграм что все предыдущие значения обработаны)
	updConf := tgbotapi.NewUpdate(0)

	// таймаут на ожидание обновлений
	updConf.Timeout = 60

	// запускаем получение апдейтов u создаем канал "updates" в который будут прилетать новые сообщения
	updates := tgBot.GetUpdatesChan(updConf)

	// в канал updates прилетают структуры типа Update - вычитываем их и обрабатываем
	for update := range updates {
		if update.Message == nil {
			continue
		}

		var reply string

		log.OutDebug(fmt.Sprintf("RECEIVED. User: %s; Chat_id: %v; Message: %s", update.Message.From.UserName, update.Message.Chat.ID, update.Message.Text))

		// прежде всего обрабатываем команды (это сообщения начинающиеся с /)
		switch update.Message.Command() {
		case "start":
			reply = "Запуск!"
		case "stop":
			log.OutInfo("Stop cmd receiving")
			return
		default:
			reply = fmt.Sprintf("А это, %s, правильный вопрос...", update.Message.From.UserName)
		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

		if _, err := tgBot.Send(msg); err != nil {
			panic(err)
		}

		// Now send message to process (to server or another)
	}

	log.Out("TgBot successfully started!")
}
