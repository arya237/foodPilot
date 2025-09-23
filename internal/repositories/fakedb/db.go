package fakedb

import (
	"sync"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
)

type FakeDb struct {
	UserMu          sync.RWMutex
	FoodMu          sync.RWMutex
	RateMu          sync.RWMutex
	Users       map[int]*models.User
	Foods       map[int]*models.Food
	Rates       map[int]map[int]*models.Rate
	UserCounter int
	FoodCounter int
}

func NewFakeDb() *FakeDb {
	return &FakeDb{
		Users:       map[int]*models.User{},
		Foods:       map[int]*models.Food{},
		Rates:       map[int]map[int]*models.Rate{},
		UserCounter: 0,
		FoodCounter: 0,
	}
}

func (fdb *FakeDb) SaveUser(username, password string) (int, error) {
	fdb.UserMu.Lock()
	defer fdb.UserMu.Unlock()
	for _, user := range fdb.Users {
		if user.Username == username {
			return 0, repositories.ErrorDuplicateUser
		}
	}

	fdb.Users[fdb.UserCounter] = &models.User{Username: username, Password: password, Id: fdb.UserCounter}
	fdb.UserCounter++
	return fdb.UserCounter - 1, nil
}

func (fdb *FakeDb) GetUserById(id int) (*models.User, error) {
	fdb.UserMu.Lock()
	defer fdb.UserMu.Unlock()
	if _, find := fdb.Users[id]; !find {
		return nil, repositories.ErrorInvalidUID
	}
	return fdb.Users[id], nil
}

func (fdb *FakeDb) GetAllUsers() ([]*models.User, error) {
	fdb.UserMu.Lock()
	defer fdb.UserMu.Unlock()
	var users []*models.User
	for _, user := range fdb.Users {
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, repositories.ErrorNoUser
	}

	return users, nil
}

func (fdb *FakeDb) DeleteUser(id int) error {
	fdb.UserMu.Lock()
	defer fdb.UserMu.Unlock()
	if _, find := fdb.Users[id]; !find {
		return repositories.ErrorInvalidUID
	}

	delete(fdb.Users, id)
	return nil
}

func (fdb *FakeDb) UpdateUser(new *models.User) error {
	fdb.UserMu.Lock()
	defer fdb.UserMu.Unlock()
	if _, find := fdb.Users[new.Id]; !find {
		return repositories.ErrorInvalidUID
	}

	fdb.Users[new.Id] = new
	return nil
}

func (fdb *FakeDb) SaveFood(name string) (int, error) {
	fdb.FoodMu.Lock()
	defer fdb.FoodMu.Unlock()
	for _, food := range fdb.Foods {
		if food.Name == name {
			return 0, repositories.ErrorDuplicateFood
		}
	}

	fdb.Foods[fdb.FoodCounter] = &models.Food{Name: name, Id: fdb.FoodCounter}
	fdb.FoodCounter++
	return fdb.FoodCounter - 1, nil
}

func (fdb *FakeDb) GetFoodById(id int) (*models.Food, error) {
	fdb.FoodMu.Lock()
	defer fdb.FoodMu.Unlock()
	if _, find := fdb.Foods[id]; !find {
		return nil, repositories.ErrorInvalidFID
	}
	return fdb.Foods[id], nil
}

func (fdb *FakeDb) GetAllFood() ([]*models.Food, error) {
	fdb.FoodMu.Lock()
	defer fdb.FoodMu.Unlock()
	var foods []*models.Food
	for _, food := range fdb.Foods {
		foods = append(foods, food)
	}

	if len(foods) == 0 {
		return nil, repositories.ErrorNoFood
	}

	return foods, nil
}

func (fdb *FakeDb) DeleteFood(id int) error {
	fdb.FoodMu.Lock()
	defer fdb.FoodMu.Unlock()
	if _, find := fdb.Foods[id]; !find {
		return repositories.ErrorInvalidFID
	}

	delete(fdb.Foods, id)
	return nil
}

func (fdb *FakeDb) UpdateFood(new *models.Food) error {
	fdb.FoodMu.Lock()
	defer fdb.FoodMu.Unlock()
	if _, find := fdb.Foods[new.Id]; !find {
		return repositories.ErrorInvalidFID
	}

	fdb.Foods[new.Id] = new
	return nil
}

func (fdb *FakeDb) SaveRate(user_id, food_id, score int) error {
	fdb.RateMu.Lock()
	defer fdb.RateMu.Unlock()
	if _, ok := fdb.Rates[user_id]; !ok {
		return repositories.ErrorInvalidUID
	}
	for _, rate := range fdb.Rates[user_id] {
		if rate.Food_id == food_id {
			return repositories.ErrorDuplicateFood
		}
	}

	if _, ok := fdb.Foods[food_id]; !ok {
		return repositories.ErrorInvalidFID
	}
	fdb.Rates[user_id][food_id] = &models.Rate{User_id: user_id, Food_id: food_id, Score: score}
	return nil
}

func (fdb *FakeDb) GetRateByUser(user_id int) ([]*models.Rate, error) {
	fdb.RateMu.Lock()
	defer fdb.RateMu.Unlock()
	if _, ok := fdb.Rates[user_id]; !ok {
		return nil, repositories.ErrorInvalidUID
	}

	var rates []*models.Rate

	for _, rate := range fdb.Rates[user_id] {
		rates = append(rates, rate)
	}
	if len(rates) == 0 {
		return nil, repositories.ErrorNorate
	}
	return rates, nil
}

func (fdb *FakeDb) DeleteRate(user_id, food_id int) error {
	fdb.RateMu.Lock()
	defer fdb.RateMu.Unlock()
	if rates, ok := fdb.Rates[user_id]; ok {
		if _, find := rates[food_id]; find {
			delete(rates, food_id)
			return nil
		} else {
			return repositories.ErrorInvalidFID
		}
	} else {
		return repositories.ErrorInvalidUID
	}
}

func (fdb *FakeDb) UpdateRate(user_id int, new *models.Rate) error {
	fdb.RateMu.Lock()
	defer fdb.RateMu.Unlock()

	fdb.RateMu.Lock()
	defer fdb.RateMu.Unlock()
	if rates, ok := fdb.Rates[user_id]; ok {
		if _, find := rates[new.Food_id]; find {
			rates[new.Food_id] = new
			return nil
		} else {
			return repositories.ErrorInvalidFID
		}
	} else {
		return repositories.ErrorInvalidUID
	}
}