package main

import (
	"awesomeAssistant/core"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	token := os.Getenv("TOKEN")
	webhookURL := os.Getenv("WEBHOOK_URL")
	localPort := os.Getenv("LOCAL_PORT")
	println(token)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhook(webhookURL + bot.Token)

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/update/" + bot.Token)
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+localPort, nil)
		if err != nil {
			log.Println(err)
		}
	}()

	for update := range updates {
		go core.HandleMessage(update, bot)
	}
}
