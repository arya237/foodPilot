package services

import (
	"errors"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/arya237/foodPilot/pkg/reservations"
	"log"
	"time"
)

type Reserve interface {
	ReserveFood() (string, error)
}

type reserve struct {
	food   FoodService
	user   UserService
	rate   RateFoodService
	samad  reservations.RequiredFunctions
	logger logger.Logger
}

func NewReserveService(f FoodService, u UserService, r RateFoodService, s reservations.RequiredFunctions) Reserve {
	return &reserve{
		food:   f,
		user:   u,
		rate:   r,
		samad:  s,
		logger: logger.New("reserve"),
	}
}

func (r *reserve) ReserveFood() (string, error) {
	users, err := r.user.GetAll()
	if err != nil {
		r.logger.Info(err.Error())
		return "", err
	}

	foods, err := r.food.GetAll()
	if err != nil {
		r.logger.Info(err.Error())
		return "", err
	}

	for _, user := range users {

		token, _ := r.samad.GetAccessToken(user.Username, user.Password)
		foodProgram, err := r.samad.GetFoodProgram(token, time.Now().Add(time.Hour*48))
		if err != nil {
			r.logger.Info(err.Error())
			return "", err
		}

		if err != nil {
			r.logger.Info(err.Error())
		}

		rates, err := r.rate.GetRateByUser(user.Id)
		if err != nil {
			r.logger.Info(err.Error())
			return "", err
		}

		for day, _ := range foodProgram.DailyFood {
			for meal, _ := range foodProgram.DailyFood[day] {
				var mealList []reservations.ReserveModel

				for _, food := range foodProgram.DailyFood[day][meal] {
					mealList = append(mealList, food)
				}

				bestFood, _ := findBestFood(mealList, rates, foods)
				log.Println(r.samad.ReserveFood(token, bestFood))
			}
		}
	}

	return "food reserved", nil
}

func findFoodWithId(foods []*models.Food, foodID int) (*models.Food, error) {
	for _, food := range foods {
		if food.Id == foodID {
			return food, nil
		}
	}

	return nil, errors.New("food not found")
}

func findBestFood(mealList []reservations.ReserveModel, rates []*models.Rate, foods []*models.Food) (reservations.ReserveModel, error) {
	bestFood := reservations.ReserveModel{}
	bestScore := 0
	for _, meal := range mealList {
		for _, rate := range rates {

			food, err := findFoodWithId(foods, rate.FoodID)
			if err != nil {
				log.Println(err.Error())
				break
			}

			if meal.FoodName == food.Name {
				if rate.Score > bestScore {
					bestFood = meal
					bestScore = rate.Score
				}
			}
		}
	}
	return bestFood, nil
}
