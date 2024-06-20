package server

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(botData botParam) {

	var msg_ch chan string

	bot, err := tgbotapi.NewBotAPI(botData.TgToken)

	if err != nil {
		fmt.Println("Error while starting the bot: ", err)
	}

	defer close(msg_ch)

	fmt.Println("Authorized on account: ", bot.Self.UserName)

	// updConf - структура с конфигом для получения апдейтов (0 - информируем телеграм что все предыдущие значения обработаны)
	updConf := tgbotapi.NewUpdate(0)

	// таймаут на ожидание обновлений
	updConf.Timeout = 60

	// запускаем получение апдейтов u создаем канал "updates" в который будут прилетать новые сообщения
	updates := bot.GetUpdatesChan(updConf)

	// в канал updates прилетают структуры типа Update - вычитываем их и обрабатываем
	for update := range updates {
		if update.Message == nil {
			continue
		}

		var reply string

		fmt.Println("bot received")
		fmt.Println("user: ", update.Message.From.UserName, "; chat_id: ", update.Message.Chat.ID, "; msg: ", update.Message.Text)

		// прежде всего обрабатываем команды (это сообщения начинающиеся с /)
		switch update.Message.Command() {
		case "start":
			reply = "Запуск!"
		case "stop":
			reply = "Стоп"
			return
		default:
			reply = fmt.Sprintf("А это, %s, правильный вопрос...", update.Message.From.UserName)
		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}

		fmt.Println("sending to main")
		msg_ch <- "Send message to user: " + update.Message.From.UserName + " by chat: " + fmt.Sprintf("%v", update.Message.Chat.ID)

		fmt.Println("sended to main")
	}
}
