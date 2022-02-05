package core

import (
	"awesomeAssistant/prometheus"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func HandleMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatId := update.Message.Chat.ID
	switch update.Message.Text {
	case "/start":
		_, err := bot.Send(tgbotapi.NewMessage(chatId, "Добро пожаловать к умному ассистенту!"))
		if err != nil {
			log.Println(err)
		}
	case "/system":
		a, b := prometheus.FreeRam()
		time := fmt.Sprint(a)
		_, err := bot.Send(tgbotapi.NewMessage(chatId, "time = "+time+", value = "+b+"%"))
		if err != nil {
			log.Println(err)
		}
	}
}
