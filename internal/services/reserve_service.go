package services

import (
	"errors"
	"sync"
	"time"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/arya237/foodPilot/pkg/reservations"
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

	const workerCount = 10
	jobs := make(chan *models.User, workerCount*2)
	var wg sync.WaitGroup

	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for user := range jobs {
				r.handleUserReservation(user, foods)
			}
		}()
	}

	go func() {
		for _, user := range users {
			if user.AutoSave {
				jobs <- user
			}
		}
		close(jobs)
	}()

	wg.Wait()
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
	logger := logger.New("findBestFoodFunction")

	for _, meal := range mealList {
		for _, rate := range rates {

			food, err := findFoodWithId(foods, rate.FoodID)
			if err != nil {
				logger.Info(err.Error())
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

func (r *reserve)handleUserReservation(user *models.User, foods []*models.Food) {
	token, _ := r.samad.GetAccessToken(user.Username, user.Password)
	foodProgram, err := r.samad.GetFoodProgram(token, time.Now().Add(time.Hour*24))

	if err != nil {
		r.logger.Info(err.Error())
	}

	// for better performance rates should be map instaed of list
	rates, err := r.rate.GetRateByUser(user.Id)
	if err != nil {
		r.logger.Info(err.Error())
	}

	for day := range foodProgram.DailyFood {
		for meal := range foodProgram.DailyFood[day] {
			mealList := foodProgram.DailyFood[day][meal]

			bestFood, _ := findBestFood(mealList, rates, foods)
			message, err := r.samad.ReserveFood(token, bestFood)
			if err != nil {
				r.logger.Info(err.Error())
			} else {
				r.logger.Info(message)
			}
		}
	}
}