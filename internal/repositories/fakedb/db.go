package fakedb

import (
	"sync"
	"github.com/arya237/foodPilot/internal/models"
)

type FakeDb struct {
	UserMu      sync.RWMutex
	FoodMu      sync.RWMutex
	RateMu      sync.RWMutex
	Users       map[int]*models.User
	Foods       map[int]*models.Food
	Rates       map[int]map[int]*models.Rate
	FoodCounter int
	UserCounter int
}

func NewDb() *FakeDb {
	return &FakeDb{
		Users:       map[int]*models.User{},
		Foods:       map[int]*models.Food{},
		Rates:       map[int]map[int]*models.Rate{},
		FoodCounter: 0,
		UserCounter: 0,
	}
}