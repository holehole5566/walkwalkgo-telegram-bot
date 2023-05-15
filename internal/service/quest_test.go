package service

import (
	"os"
	db "telebot/internal/access"
	"testing"
)

func TestGetQuestWithFinishStatus(t *testing.T) {
	mongoURI := os.Getenv("MONGO_URI")
	client := db.GetClient(mongoURI)
	arrivals := GetArrivals(1120218288, 2023, 18, client)
	quest := GetQuest(2023, 18, client)
	result := GetQuestWithFinishStatus(quest, arrivals)
	t.Log(result)
}
