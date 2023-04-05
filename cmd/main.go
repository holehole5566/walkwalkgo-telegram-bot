package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	row      []tgbotapi.KeyboardButton
	keyboard tgbotapi.ReplyKeyboardMarkup
)

func init() {

	btn := tgbotapi.KeyboardButton{
		RequestLocation: true,
		Text:            "Arrived!",
	}

	row = tgbotapi.NewKeyboardButtonRow(btn)
	keyboard = tgbotapi.NewReplyKeyboard(row)
}

func main() {

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("Token must be set")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			if update.Message.Location != nil {

				log.Printf("Location：%f,%f", update.Message.Location.Latitude, update.Message.Location.Longitude)

				reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Good job!")
				if _, err := bot.Send(reply); err != nil {
					log.Panic(err)
				}
			} else {
				switch update.Message.Text {
				case "/start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, welcome to WalkWalk Go.")
					msg.ReplyMarkup = keyboard
					bot.Send(msg)
				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please use keyboard button to send your location.")
					msg.ReplyMarkup = keyboard
					bot.Send(msg)
				}
			}
		}
	}
}
