package memory

import (
	"github.com/arya237/foodPilot/internal/db/tempdb"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
)

type rateRepo struct {
	db *tempdb.FakeDb
}

func NewRateRepo(db *tempdb.FakeDb) repositories.Rate {
	return &rateRepo{
		db: db,
	}
}

func (fdb *rateRepo) Save(userID, foodID, score int) error {
	fdb.db.RateMu.Lock()
	defer fdb.db.RateMu.Unlock()
	if _, ok := fdb.db.Users[userID]; !ok {
		return repositories.ErrorNotFound
	}

	if _, ok := fdb.db.Foods[foodID]; !ok {
		return repositories.ErrorNotFound
	}
	if fdb.db.Rates[userID] == nil {
		fdb.db.Rates[userID] = make(map[int]*models.Rate)
	}
	fdb.db.Rates[userID][foodID] = &models.Rate{UserID: userID, FoodID: foodID, Score: score}
	return nil
}

func (fdb *rateRepo) GetByUser(userID int) ([]*models.Rate, error) {
	fdb.db.RateMu.RLock()
	defer fdb.db.RateMu.RUnlock()
	if _, ok := fdb.db.Rates[userID]; !ok {
		return nil, nil
	}

	var rates []*models.Rate

	for _, rate := range fdb.db.Rates[userID] {
		rates = append(rates, rate)
	}

	return rates, nil
}

func (fdb *rateRepo) Delete(userID, foodID int) error {
	fdb.db.RateMu.Lock()
	defer fdb.db.RateMu.Unlock()

	if rates, ok := fdb.db.Rates[userID]; ok {
		if _, find := rates[foodID]; find {
			delete(rates, foodID)
			return nil
		} else {
			return repositories.ErrorNotFound
		}
	}
	return repositories.ErrorNotFound
}

func (fdb *rateRepo) Update(userID int, new *models.Rate) error {
	fdb.db.RateMu.Lock()
	defer fdb.db.RateMu.Unlock()

	if rates, ok := fdb.db.Rates[userID]; ok {
		if _, find := rates[new.FoodID]; find {
			rates[new.FoodID] = new
			return nil
		} else {
			return repositories.ErrorNotFound
		}
	}
	return repositories.ErrorNotFound
}
