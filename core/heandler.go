package core

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func HandleMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatId := update.Message.Chat.ID
	println(update.Message.Text)
	_, err := bot.Send(tgbotapi.NewMessage(chatId, "Hello"))
	if err != nil {
		log.Println(err)
	}

}
