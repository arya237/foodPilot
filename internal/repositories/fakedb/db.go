package fakedb

import (
	"github.com/arya237/foodPilot/internal/models"
	"sync"
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
		Users: map[int]*models.User{},
		Foods: map[int]*models.Food{1: {Name: "چلو کباب کوبیده زعفرانی", Id: 1},
			2: {Name: "چلو جوجه کباب", Id: 2}, 3: {Name: "خوراک گوشت چرخ‌کرده با سیب زمینی", Id: 3}},
		Rates:       map[int]map[int]*models.Rate{},
		FoodCounter: 0,
		UserCounter: 0,
	}
}
