package main

import (
	"context"
	"log"
	"net/http"
	"os"
	db "telebot/internal/access"
	"telebot/internal/service"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	row      []tgbotapi.KeyboardButton
	keyboard tgbotapi.ReplyKeyboardMarkup
	reply    tgbotapi.MessageConfig
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

	go func() {
		port := ":8080"
		handler := http.FileServer(http.Dir("web"))
		http.ListenAndServe(port, handler)

	}()

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("TOKEN must be set")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI must be set")
	}

	client := db.GetClient(mongoURI)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	log.Printf("Pinged your deployment. You successfully connected to MongoDB!")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	log.Printf("Start to listen...")

	for update := range updates {
		if update.Message != nil {

			if update.Message.Location != nil {

				log.Printf("Location：%f,%f", update.Message.Location.Latitude, update.Message.Location.Longitude)

				fix := time.FixedZone("UTC+8", 3600*8)
				tm := time.Unix(int64(update.Message.Date), 0)
				localTime := tm.In(fix)
				checkInResult := ""
				result := service.CheckArrivedStatus(update.Message.Chat.ID, localTime, client, update.Message.Location.Latitude, update.Message.Location.Longitude)

				switch result {
				case 1:
					checkInResult = "Check in Success. Good Job!"
				case 2:
					checkInResult = "You are not in a spot area."
				case 3:
					checkInResult = "You have already checked in this spot."
				case 4:
					checkInResult = "Something went wrong, please contact Peanutbutter."
				default:
					checkInResult = "Something went wrong, please try again."
				}

				reply = tgbotapi.NewMessage(update.Message.Chat.ID, checkInResult)
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
