package service

import (
	"fmt"
	"telebot/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetQuestWithFinishStatus(quest model.Quest, arrivals []model.Arrival) model.Quest {
	spotWithFinished := []model.Spot{}
	questSpot := quest.Spots
	for _, spot := range questSpot {
		spot.Finished = false
		for _, arrival := range arrivals {
			if spot.Name == arrival.Spot.Name {
				spot.Finished = true
			}
		}
		spotWithFinished = append(spotWithFinished, spot)
	}
	result := model.Quest{
		Desc:  quest.Desc,
		Year:  quest.Year,
		Week:  quest.Week,
		Spots: spotWithFinished,
	}
	return result
}

func GetQuestResult(userID int64, dateTime time.Time, client *mongo.Client) string {
	//year, week := dateTime.ISOWeek()
	year := 2023
	week := 18
	arrivals := GetArrivals(int(userID), year, week, client)
	quest := GetQuest(year, week, client)
	questResult := GetQuestWithFinishStatus(quest, arrivals)
	spot1 := questResult.Spots[0]
	spot2 := questResult.Spots[1]
	spot3 := questResult.Spots[2]
	result := fmt.Sprintf(`
<b>Year %d, Week %d Quest</b>
%s
<b>Spot 1</b>
Name: %s
Location: %s
Finished: %t
<b>Spot 2</b>
Name: %s
Location: %s
Finished: %t
<b>Spot 3</b>
Name: %s
Location: %s
Finished: %t
`, year, week, questResult.Desc, spot1.Name, GetLocationURL(spot1.Latitude, spot1.Longitude), spot1.Finished, spot2.Name, GetLocationURL(spot2.Latitude, spot2.Longitude), spot2.Finished, spot3.Name, GetLocationURL(spot3.Latitude, spot3.Longitude), spot3.Finished)
	return result
}

func GetLocationURL(latitude float64, longitude float64) string {
	lati := fmt.Sprintf("%f", latitude)
	long := fmt.Sprintf("%f", longitude)
	return fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%s,%s", lati, long)
}
