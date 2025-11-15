package services

import (
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
)

type RateFoodService interface {
	GetRateByUser(userID int) (map[string]int, error)
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

func (s *rateFoodService) GetRateByUser(userID int) (map[string]int, error) {
	rates, err := s.RateRepo.GetByUser(userID)
	if err != nil {
		return nil, err
	}

	userRates := make(map[string]int, len(rates))
	for _, rate := range rates {
		food, _ := s.FoodRepo.GetById(rate.FoodID)
		userRates[food.Name] = rate.Score
	}

	return userRates, nil
}
