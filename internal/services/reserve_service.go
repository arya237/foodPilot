package services

import (
	"fmt"
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
	user   UserService
	rate   RateFoodService
	samad  reservations.RequiredFunctions
	logger logger.Logger
}

func NewReserveService(u UserService, r RateFoodService, s reservations.RequiredFunctions) Reserve {
	return &reserve{
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

	const workerCount = 10
	jobs := make(chan *models.User, workerCount*2)
	var wg sync.WaitGroup

	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for user := range jobs {
				mmap, err := r.handleUserReservation(user)
				r.logger.Info("geting user info -> "+user.Username,
					logger.Field{
						Key:   "map",
						Value: mmap,
					},
					logger.Field{
						Key:   "error",
						Value: err,
					},
				)
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

func findBestFood(mealList []reservations.ReserveModel, rates map[string]int) (reservations.ReserveModel, error) {
	bestFood := reservations.ReserveModel{}
	bestScore := 0

	for _, meal := range mealList {
		for foodName, rate := range rates {
			if meal.FoodName == foodName {
				if rate > bestScore {
					bestFood = meal
					bestScore = rate
				}
			}
		}
	}
	return bestFood, nil
}

func (r *reserve) handleUserReservation(user *models.User) (reservation map[string]map[string]string, err error) {
	// TODO: check if token is valid or not
	token, _ := r.samad.GetAccessToken(user.Username, user.Password)

	// Get Samad food program
	foodProgram, err := r.samad.GetFoodProgram(token, time.Now().Add(time.Hour*24))
	if err != nil {
		r.logger.Info(err.Error())
		return nil, err
	}

	if foodProgram == nil {
		r.logger.Warn("this user food program is nil",
			logger.Field{Key: "User", Value: user},
		)
		return nil, fmt.Errorf("user %s has get nil from samad", user.Username)
	}

	// Get user rates
	rates, err := r.rate.GetRateByUser(user.Id)
	if err != nil {
		r.logger.Info(err.Error())
		return nil, err
	}

	// collect per-day reservation errors while continuing other days
	dayErrors := make(map[string]map[string]string)
	for day := range foodProgram.DailyFood {
		for meal := range foodProgram.DailyFood[day] {

			mealList := foodProgram.DailyFood[day][meal]
			bestFood, _ := findBestFood(mealList, rates)
			message, err := r.samad.ReserveFood(token, bestFood)

			if err != nil {
				message = err.Error()
			}

			dayErrors[day.String()][meal.String()] = message
		}
	}

	return dayErrors, nil
}
