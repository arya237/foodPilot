package services

import (
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/arya237/foodPilot/internal/getways/reservations"
)

type Reserve interface {
	UserReservation(userID int) (*UserReserveResult, error)
	ReserveFood() ([]*UserReserveResult, error)
}

//**********************    Structured reservation results    ***********************************

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
	Error    string      `json:"error,omitempty"`
	Days     []DayResult `json:"days,omitempty"`
}

const (
	heroFood = "املت"
)

//******************************************************************************

type reserve struct {
	user       repositories.User
	userCred   repositories.RestaurantCredentials
	userSevise UserService
	samad      reservations.ReserveFunctions
	logger     logger.Logger
}

func NewReserveService(u repositories.User, userCred repositories.RestaurantCredentials, userService UserService, s reservations.ReserveFunctions) Reserve {
	return &reserve{
		user:       u,
		userCred:   userCred,
		userSevise: userService,
		samad:      s,
		logger:     logger.New("reserve"),
	}
}

func (r *reserve) ReserveFood() ([]*UserReserveResult, error) {
	users, err := r.user.GetAll()
	if err != nil {
		r.logger.Info(err.Error())
		return nil, err
	}

	const workerCount = 10
	jobs := make(chan *models.User, workerCount*2)
	results := make(chan *UserReserveResult, workerCount*2)
	var wg sync.WaitGroup

	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for user := range jobs {
				res, err := r.UserReservation(user.Id)
				if err != nil {
					r.logger.Warn(err.Error())
				}
				results <- res
			}
		}()
	}

	// go func() {
	// 	for _, user := range users {
	// 		if user.AutoSave {
	// 			jobs <- user
	// 		}
	// 	}
	// 	close(jobs)
	// }()

	// collect results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Add all ansers togheter
	aggregated := make([]*UserReserveResult, 0, len(users))
	for res := range results {
		aggregated = append(aggregated, res)
	}

	return aggregated, nil
}

func findBestFood(mealList []reservations.ReserveModel, rates map[string]int) (reservations.ReserveModel, error) {
	bestFood := mealList[0]
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

func (r *reserve) UserReservation(userID int) (*UserReserveResult, error) {
	// TODO: check if token is valid or not

	cred, err := r.userCred.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	token, err := r.samad.GetAccessToken(cred.Username, cred.Password)
	if err != nil {
		return nil, ErrTokenGeneration
	}

	selfIDs, err := r.samad.GetProperSelfID(token)
	if err != nil {
		r.logger.Info(err.Error())
		return nil, err
	}

	selfsFoodProgram := make(map[string]*reservations.WeekFood)

	for selfName, selfID := range selfIDs {
		selfsFoodProgram[selfName], err = r.samad.GetFoodProgram(token, selfID, time.Now().Add(time.Hour*48))
		if err != nil {
			r.logger.Info(err.Error())
			return &UserReserveResult{UserID: cred.UserID, Username: cred.Username, Error: err.Error()}, err
		}
	}

	// Get user rates
	rates, err := r.userSevise.ViewRating(cred.UserID)
	if err != nil {
		r.logger.Info(err.Error())
		return &UserReserveResult{UserID: cred.UserID, Username: cred.Username, Error: err.Error()}, err
	}

	// build structured per-day results while continuing on errors
	dayResults := make([]DayResult, 0, 7 /*food count*/)
	//log.Print(foodProgram.DailyFood)
	for key, value := range selfsFoodProgram {
		if strings.Contains(key, "مرکزی") {
			for day := range value.DailyFood {
				meals := make([]MealResult, 0, 3 /*meals count*/)

				for meal := range value.DailyFood[day] {
					mealList := value.DailyFood[day][meal]
					if heroMeal, err := lookingForSpecificFood(heroFood, selfsFoodProgram, day, meal); err == nil {
						mealList = append([]reservations.ReserveModel{*heroMeal}, mealList...)
					} else {
						log.Printf("\n\n\n\n%s\n\n\n", err.Error())
					}
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

		}
	}

	return &UserReserveResult{UserID: cred.UserID, Username: cred.Username, Days: dayResults}, nil
}

func lookingForSpecificFood(foodName string, foodprogram map[string]*reservations.WeekFood, day reservations.Weekday, meal reservations.Meal) (*reservations.ReserveModel, error) {
	for _, value := range foodprogram {
		if _, exist := value.DailyFood[day]; !exist {
			return nil, errors.New("day " + day.String() + " doesn't exist")
		}
		if _, exist := value.DailyFood[day][meal]; !exist {
			return nil, errors.New("meal " + meal.String() + " doesn't exist")
		}

		for _, food := range value.DailyFood[day][meal] {
			if food.FoodName == foodName {

				return &food, nil
			}
		}
	}

	return nil, errors.New("food " + foodName + " doesn't exist")
}
