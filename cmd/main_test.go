package main

import (
	"os"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {

	btn := tgbotapi.KeyboardButton{
		RequestLocation: true,
		Text:            "Arrived!",
	}

	row = tgbotapi.NewKeyboardButtonRow(btn)
	keyboard = tgbotapi.NewReplyKeyboard(row)
}

func TestInitKeyboard(t *testing.T) {
	expectedText := "Arrived!"
	if len(keyboard.Keyboard) != 1 {
		t.Error("Expected one row of keyboard buttons")
	}
	row := keyboard.Keyboard[0]
	if len(row) != 1 {
		t.Error("Expected one button in the row")
	}
	btn := row[0]
	if btn.Text != expectedText {
		t.Errorf("Expected button text to be %s but got %s", expectedText, btn.Text)
	}
	if !btn.RequestLocation {
		t.Error("Expected button to request location")
	}
}

func TestGetenv(t *testing.T) {
	token := os.Getenv("TOKEN")
	if token == "" {
		t.Fatal("Token must be set")
	}
}
