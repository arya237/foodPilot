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
	ReserveFood() ([]UserReserveResult, error)
}

// Structured reservation results
type MealResult struct {
	Meal    reservations.Meal `json:"meal"`
	Message string            `json:"message"`
	Ok      bool              `json:"ok"`
}

type DayResult struct {
	Day   reservations.Weekday `json:"day"`
	Meals []MealResult         `json:"meals,omitempty"`
}

type UserReserveResult struct {
	UserID   int         `json:"user_id"`
	Username string      `json:"username"`
	Days     []DayResult `json:"days,omitempty"`
}
//******************************************************************************

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

func (r *reserve) ReserveFood() ([]UserReserveResult, error) {
	users, err := r.user.GetAll()
	if err != nil {
		r.logger.Info(err.Error())
		return nil, err
	}

	const workerCount = 10
	jobs := make(chan *models.User, workerCount*2)
	results := make(chan UserReserveResult, workerCount*2)
	var wg sync.WaitGroup

	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for user := range jobs {
				res, err := r.handleUserReservation(user)
				if err != nil {
					r.logger.Info(err.Error())
				}
				results <- res
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

	// collect results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Add all ansers togheter
	aggregated := make([]UserReserveResult, 0, len(users))
	for res := range results {
		aggregated = append(aggregated, res)
	}

	return aggregated, nil
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

func (r *reserve) handleUserReservation(user *models.User) (UserReserveResult, error) {
	// TODO: check if token is valid or not
	token, _ := r.samad.GetAccessToken(user.Username, user.Password)

	// Get Samad food program
	foodProgram, err := r.samad.GetFoodProgram(token, time.Now().Add(time.Hour*24))
	if err != nil {
		r.logger.Info(err.Error())
		return UserReserveResult{UserID: user.Id, Username: user.Username}, err
	}

	if foodProgram == nil {
		r.logger.Warn("this user food program is nil",
			logger.Field{Key: "User", Value: user},
		)
		return UserReserveResult{UserID: user.Id, Username: user.Username}, fmt.Errorf("user %s has get nil from samad", user.Username)
	}

	// Get user rates
	rates, err := r.rate.GetRateByUser(user.Id)
	if err != nil {
		r.logger.Info(err.Error())
		return UserReserveResult{UserID: user.Id, Username: user.Username}, err
	}

	// build structured per-day results while continuing on errors
	dayResults := make([]DayResult, 0, 7/*food count*/)
	for day := range foodProgram.DailyFood {
		meals := make([]MealResult, 0, 3/*meals count*/)

		for meal := range foodProgram.DailyFood[day] {
			mealList := foodProgram.DailyFood[day][meal]
			bestFood, _ := findBestFood(mealList, rates)
			message, err := r.samad.ReserveFood(token, bestFood)

			if err != nil {
				meals = append(meals, MealResult{Meal: meal, Message: err.Error(), Ok: false})
				continue
			}

			meals = append(meals, MealResult{Meal: meal, Message: message, Ok: true})
		}

		dayResults = append(dayResults, DayResult{Day: day, Meals: meals})
	}

	return UserReserveResult{UserID: user.Id, Username: user.Username, Days: dayResults}, nil
}
