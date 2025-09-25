package services

import (
	"errors"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
	"strconv"
)

type RateFoodService interface {
	SaveRate(userID string, foods map[string]int) (string, error)
	GetRateByUser(userID int) ([]*models.Rate, error)
}

type rateFoodService struct {
	FoodRepo repositories.Food
	RateRepo repositories.Rate
	logger   logger.Logger
}

func NewRateFoodService(r repositories.Rate, f repositories.Food) RateFoodService {
	return &rateFoodService{
		FoodRepo: f,
		RateRepo: r,
		logger:   logger.New("rateFood_service"),
	}
}

func (s *rateFoodService) SaveRate(userID string, foods map[string]int) (string, error) {

	foodList, err := s.FoodRepo.GetAllFood()

	if err != nil {
		s.logger.Info(err.Error())
		return "", err
	}

	for key, value := range foods {

		foodID, err := findFoodID(foodList, key)

		if err != nil {
			s.logger.Info(err.Error())
			return "", err
		}

		userID, err := strconv.Atoi(userID)
		if err != nil {
			s.logger.Info(err.Error())
		}
		err = s.RateRepo.SaveRate(userID, foodID, value)

		if err != nil {
			s.logger.Info(err.Error())
			return "", err
		}
	}

	return "all Rates save successfully", nil
}

func (s *rateFoodService) GetRateByUser(userID int) ([]*models.Rate, error) {
	return s.RateRepo.GetRateByUser(userID)
}

func findFoodID(foods []*models.Food, foodName string) (int, error) {
	for _, food := range foods {
		if food.Name == foodName {
			return food.Id, nil
		}
	}

	return 0, errors.New("food not found")
}
