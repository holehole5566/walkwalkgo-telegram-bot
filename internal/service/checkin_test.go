package service

import (
	"log"
	"os"
	db "telebot/internal/access"
	model "telebot/internal/model"
	"testing"
)

func TestGetQuest(t *testing.T) {
	mongoURI := os.Getenv("MONGO_URI")
	client := db.GetClient(mongoURI)
	Quest := GetQuest(2023, 18, client)
	log.Println(Quest)
}

func TestGetArrivals(t *testing.T) {
	mongoURI := os.Getenv("MONGO_URI")
	client := db.GetClient(mongoURI)
	Arrivals := GetArrivals(1120218288, 2023, 18, client)
	log.Println(Arrivals)
}

func TestCheckQuestArrival(t *testing.T) {
	mongoURI := os.Getenv("MONGO_URI")
	client := db.GetClient(mongoURI)
	Quest := GetQuest(2023, 18, client)
	spot := CheckQuestArrival(25.0339639, 121.5644722, Quest)
	if spot.Name == "" {
		log.Println("Not in the quest area")
	}
	log.Println(spot)

}

func TestAlreadyArrived(t *testing.T) {
	mongoURI := os.Getenv("MONGO_URI")
	client := db.GetClient(mongoURI)
	Quest := GetQuest(2023, 18, client)
	spot := CheckQuestArrival(25.0339639, 121.5644722, Quest)
	Arrivals := GetArrivals(1120218288, 2023, 18, client)
	isAlreadyArrived := CheckAlreadyArrived(spot, Arrivals)
	log.Println(isAlreadyArrived)
}

func TestInsertArrival(t *testing.T) {
	mongoURI := os.Getenv("MONGO_URI")
	client := db.GetClient(mongoURI)
	spot := model.Spot{
		Name:      "Test Spot",
		Type:      1,
		Latitude:  25.55667788,
		Longitude: 121.55667788,
	}
	_, err := InsertArrival(int(1120218288), 2023, 18, spot, client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("insert success")
	}
}
