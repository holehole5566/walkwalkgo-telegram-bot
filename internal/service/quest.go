package service

import (
	"telebot/internal/model"
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
