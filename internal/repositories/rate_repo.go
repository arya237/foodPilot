package repositories

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories/fakedb"
)

type Rate interface {
	SaveRate(user_id, food_id, score int) error
	GetRateByUser(user_id int) ([]*models.Rate, error)
	DeleteRate(user_id, rate_id int) error
	UpdateRate(user_id int, new *models.Rate) error
}

type rateRepo struct {
	db *fakedb.FakeDb
}

func NewRateRepo(db *fakedb.FakeDb) Rate {
	return &rateRepo{
		db: db,
	}
}

func (fdb *rateRepo) SaveRate(user_id, food_id, score int) error {
	fdb.db.RateMu.Lock()
	defer fdb.db.RateMu.Unlock()
	if _, ok := fdb.db.Rates[user_id]; !ok {
		return ErrorInvalidUID
	}
	for _, rate := range fdb.db.Rates[user_id] {
		if rate.Food_id == food_id {
			return ErrorDuplicateFood
		}
	}

	if _, ok := fdb.db.Foods[food_id]; !ok {
		return ErrorInvalidFID
	}
	fdb.db.Rates[user_id][food_id] = &models.Rate{User_id: user_id, Food_id: food_id, Score: score}
	return nil
}

func (fdb *rateRepo) GetRateByUser(user_id int) ([]*models.Rate, error) {
	fdb.db.RateMu.RLock()
	defer fdb.db.RateMu.Unlock()
	if _, ok := fdb.db.Rates[user_id]; !ok {
		return nil, ErrorInvalidUID
	}

	var rates []*models.Rate

	for _, rate := range fdb.db.Rates[user_id] {
		rates = append(rates, rate)
	}
	if len(rates) == 0 {
		return nil, ErrorNorate
	}
	return rates, nil
}

func (fdb *rateRepo) DeleteRate(user_id, food_id int) error {
	fdb.db.RateMu.Lock()
	defer fdb.db.RateMu.Unlock()
	if rates, ok := fdb.db.Rates[user_id]; ok {
		if _, find := rates[food_id]; find {
			delete(rates, food_id)
			return nil
		} else {
			return ErrorInvalidFID
		}
	} else {
		return ErrorInvalidUID
	}
}

func (fdb *rateRepo) UpdateRate(user_id int, new *models.Rate) error {
	fdb.db.RateMu.Lock()
	defer fdb.db.RateMu.Unlock()

	fdb.db.RateMu.Lock()
	defer fdb.db.RateMu.Unlock()
	if rates, ok := fdb.db.Rates[user_id]; ok {
		if _, find := rates[new.Food_id]; find {
			rates[new.Food_id] = new
			return nil
		} else {
			return ErrorInvalidFID
		}
	} else {
		return ErrorInvalidUID
	}
}