package repositories

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories/fakedb"
)

type Rate interface {
	SaveRate(userID, foodID, score int) error
	GetRateByUser(userID int) ([]*models.Rate, error)
	DeleteRate(userID, foodID int) error
	UpdateRate(userID int, new *models.Rate) error
}

type rateRepo struct {
	db *fakedb.FakeDb
}

func NewRateRepo(db *fakedb.FakeDb) Rate {
	return &rateRepo{
		db: db,
	}
}

func (fdb *rateRepo) SaveRate(userID, foodID, score int) error {
	fdb.db.RateMu.Lock()
	defer fdb.db.RateMu.Unlock()
	if _, ok := fdb.db.Users[userID]; !ok {
		return ErrorInvalidUID
	}
	for _, rate := range fdb.db.Rates[userID] {
		if rate.FoodID == foodID {
			return ErrorDuplicateFood
		}
	}

	if _, ok := fdb.db.Foods[foodID]; !ok {
		return ErrorInvalidFID
	}
	if fdb.db.Rates[userID] == nil {
		fdb.db.Rates[userID] = make(map[int]*models.Rate)
	}
	fdb.db.Rates[userID][foodID] = &models.Rate{UserID: userID, FoodID: foodID, Score: score}
	return nil
}

func (fdb *rateRepo) GetRateByUser(userID int) ([]*models.Rate, error) {
	fdb.db.RateMu.RLock()
	defer fdb.db.RateMu.RUnlock()
	if _, ok := fdb.db.Rates[userID]; !ok {
		return nil, ErrorInvalidUID
	}

	var rates []*models.Rate

	for _, rate := range fdb.db.Rates[userID] {
		rates = append(rates, rate)
	}
	if len(rates) == 0 {
		return nil, ErrorNorate
	}
	return rates, nil
}

func (fdb *rateRepo) DeleteRate(userID, foodID int) error {
	fdb.db.RateMu.Lock()
	defer fdb.db.RateMu.Unlock()
	if rates, ok := fdb.db.Rates[userID]; ok {
		if _, find := rates[foodID]; find {
			delete(rates, foodID)
			return nil
		} else {
			return ErrorInvalidFID
		}
	} else {
		return ErrorInvalidUID
	}
}

func (fdb *rateRepo) UpdateRate(userID int, new *models.Rate) error {
	fdb.db.RateMu.Lock()
	defer fdb.db.RateMu.Unlock()

	if rates, ok := fdb.db.Rates[userID]; ok {
		if _, find := rates[new.FoodID]; find {
			rates[new.FoodID] = new
			return nil
		} else {
			return ErrorInvalidFID
		}
	} else {
		return ErrorInvalidUID
	}
}
