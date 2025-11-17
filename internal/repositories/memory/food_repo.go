package memory

import (
	"github.com/arya237/foodPilot/internal/db"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
)

type foodRepo struct {
	db *db.FakeDb
}

func NewFoodRepo(db *db.FakeDb) repositories.Food {
	return &foodRepo{
		db: db,
	}
}

func (fdb *foodRepo) Save(name string) (int, error) {
	fdb.db.FoodMu.Lock()
	defer fdb.db.FoodMu.Unlock()

	for _, food := range fdb.db.Foods {
		if food.Name == name {
			return 0, repositories.ErrorDuplicateFood
		}
	}

	fdb.db.Foods[fdb.db.FoodCounter] = &models.Food{Name: name, Id: fdb.db.FoodCounter}
	fdb.db.FoodCounter++
	return fdb.db.FoodCounter - 1, nil
}

func (fdb *foodRepo) GetById(id int) (*models.Food, error) {
	fdb.db.FoodMu.RLock()
	defer fdb.db.FoodMu.RUnlock()
	if _, find := fdb.db.Foods[id]; !find {
		return nil, repositories.ErrorInvalidFID
	}
	return fdb.db.Foods[id], nil
}

func (fdb *foodRepo) GetAll() ([]*models.Food, error) {
	fdb.db.FoodMu.RLock()
	defer fdb.db.FoodMu.RUnlock()
	var foods []*models.Food
	for _, food := range fdb.db.Foods {
		foods = append(foods, food)
	}

	if len(foods) == 0 {
		return nil, repositories.ErrorNoFood
	}

	return foods, nil
}

func (fdb *foodRepo) Delete(id int) error {
	fdb.db.FoodMu.Lock()
	defer fdb.db.FoodMu.Unlock()
	if _, find := fdb.db.Foods[id]; !find {
		return repositories.ErrorInvalidFID
	}

	delete(fdb.db.Foods, id)
	return nil
}

func (fdb *foodRepo) Update(new *models.Food) error {
	fdb.db.FoodMu.Lock()
	defer fdb.db.FoodMu.Unlock()
	if _, find := fdb.db.Foods[new.Id]; !find {
		return repositories.ErrorInvalidFID
	}

	fdb.db.Foods[new.Id] = new
	return nil
}
