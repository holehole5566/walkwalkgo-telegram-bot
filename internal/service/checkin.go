package service

import (
	"context"
	"log"
	"math"
	"telebot/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// return 1: arrived  2: not in the quest area 3: already arrived 4: error
func CheckArrivedStatus(userID int64, dateTime time.Time, client *mongo.Client, lantitude float64, longtitude float64) int {

	year, week := dateTime.ISOWeek()
	quest := GetQuest(year, week, client)
	arrivals := GetArrivals(int(userID), year, week, client)

	spot := CheckQuestArrival(lantitude, longtitude, quest)
	if spot.Name == "" {
		return 2
	}
	isAlreadyArrived := CheckAlreadyArrived(spot, arrivals)
	if isAlreadyArrived {
		return 3
	}
	_, err := InsertArrival(int(userID), year, week, spot, client)
	if err != nil {
		return 4
	}

	return 1
}

func GetQuest(year int, week int, client *mongo.Client) model.Quest {

	collection := client.Database("walkwalkgo").Collection("quest")
	var results []model.Quest
	filter := bson.D{{"week", week}, {"year", year}}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}
	return results[0]
}

func GetArrivals(user_ID int, year int, week int, client *mongo.Client) []model.Arrival {
	collection := client.Database("walkwalkgo").Collection("arrival")
	var results []model.Arrival
	filter := bson.D{{"week", week}, {"year", year}, {"user_id", user_ID}}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}
	return results
}

func CheckQuestArrival(lantitude float64, longtitude float64, quest model.Quest) model.Spot {
	result := model.Spot{}
	for _, spot := range quest.Spots {
		if CheckDistance(lantitude, longtitude, spot.Latitude, spot.Longitude) {
			result = spot
		}
	}
	return result
}

func CheckDistance(latitude1 float64, longitude1 float64, lantitude2 float64, longitude2 float64) bool {
	distance := CalculateDistance(latitude1, longitude1, lantitude2, longitude2)

	return distance <= 0.05
}

func CalculateDistance(latitude1 float64, longitude1 float64, lantitude2 float64, longitude2 float64) float64 {

	lat1 := degreeToRadian(latitude1)
	lon1 := degreeToRadian(longitude1)
	lat2 := degreeToRadian(lantitude2)
	lon2 := degreeToRadian(longitude2)

	dLat := lat2 - lat1
	dLon := lon2 - lon1
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := 6371 * c

	return d
}

func degreeToRadian(deg float64) float64 {
	return deg * math.Pi / 180
}

func CheckAlreadyArrived(spot model.Spot, arrivals []model.Arrival) bool {
	result := false
	for _, arrival := range arrivals {
		if arrival.Spot.Name == spot.Name {
			result = true
		}
	}
	return result
}

func InsertArrival(userID int, year int, week int, spot model.Spot, client *mongo.Client) (bool, error) {
	collection := client.Database("walkwalkgo").Collection("arrival")
	arrival := model.Arrival{
		UserID: userID,
		Year:   year,
		Week:   week,
		Spot:   spot,
	}
	_, err := collection.InsertOne(context.TODO(), arrival)
	if err != nil {
		return false, err
	}

	return true, err
}
