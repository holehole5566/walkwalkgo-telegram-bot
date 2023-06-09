package main

import (
	"context"
	"log"
	"os"
	db "telebot/internal/access"
	"telebot/internal/service"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	row       []tgbotapi.KeyboardButton
	keyboard  tgbotapi.ReplyKeyboardMarkup
	reply     tgbotapi.MessageConfig
	keyboard2 tgbotapi.ReplyKeyboardMarkup
)

func init() {

	btn := tgbotapi.KeyboardButton{
		RequestLocation: true,
		Text:            "Arrived!",
	}

	row = tgbotapi.NewKeyboardButtonRow(btn)
	keyboard = tgbotapi.NewReplyKeyboard(row)

	keyboard2 = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/checkin"),
			tgbotapi.NewKeyboardButton("/profile"),
			tgbotapi.NewKeyboardButton("/quests"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/achievements"),
		),
	)
}

func main() {

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

			fix := time.FixedZone("UTC+8", 3600*8)
			tm := time.Unix(int64(update.Message.Date), 0)
			localTime := tm.In(fix)

			if update.Message.Location != nil {

				log.Printf("Location：%f,%f", update.Message.Location.Latitude, update.Message.Location.Longitude)

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
					msg.ReplyMarkup = keyboard2
					bot.Send(msg)
				case "/checkin":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, welcome to WalkWalk Go.")
					msg.ReplyMarkup = keyboard
					bot.Send(msg)
				case "/profile":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					msg.ParseMode = tgbotapi.ModeHTML
					bot.Send(msg)
				case "/quests":
					result := service.GetQuestResult(update.Message.Chat.ID, localTime, client)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
					msg.ParseMode = tgbotapi.ModeHTML
					bot.Send(msg)
				case "/achievements":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					msg.ParseMode = tgbotapi.ModeHTML
					bot.Send(msg)
				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Select the button below.")
					msg.ReplyMarkup = keyboard2
					bot.Send(msg)
				}
			}
		}
	}
}
