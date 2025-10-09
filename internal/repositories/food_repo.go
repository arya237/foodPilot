package repositories

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories/fakedb"
)

type Food interface {
	SaveFood(name string) (int, error)
	GetFoodById(id int) (*models.Food, error)
	GetAllFood() ([]*models.Food, error)
	DeleteFood(id int) error
	UpdateFood(new *models.Food) error
}

type foodRepo struct {
	db *fakedb.FakeDb
}

func NewFoodRepo(db *fakedb.FakeDb) Food {
	return &foodRepo{
		db: db,
	}
}

func (fdb *foodRepo) SaveFood(name string) (int, error) {
	fdb.db.FoodMu.Lock()
	defer fdb.db.FoodMu.Unlock()

	for _, food := range fdb.db.Foods {
		if food.Name == name {
			return 0, ErrorDuplicateFood
		}
	}

	fdb.db.Foods[fdb.db.FoodCounter] = &models.Food{Name: name, Id: fdb.db.FoodCounter}
	fdb.db.FoodCounter++
	return fdb.db.FoodCounter - 1, nil
}

func (fdb *foodRepo) GetFoodById(id int) (*models.Food, error) {
	fdb.db.FoodMu.RLock()
	defer fdb.db.FoodMu.RUnlock()
	if _, find := fdb.db.Foods[id]; !find {
		return nil, ErrorInvalidFID
	}
	return fdb.db.Foods[id], nil
}

func (fdb *foodRepo) GetAllFood() ([]*models.Food, error) {
	fdb.db.FoodMu.RLock()
	defer fdb.db.FoodMu.RUnlock()
	var foods []*models.Food
	for _, food := range fdb.db.Foods {
		foods = append(foods, food)
	}

	if len(foods) == 0 {
		return nil, ErrorNoFood
	}

	return foods, nil
}

func (fdb *foodRepo) DeleteFood(id int) error {
	fdb.db.FoodMu.Lock()
	defer fdb.db.FoodMu.Unlock()
	if _, find := fdb.db.Foods[id]; !find {
		return ErrorInvalidFID
	}

	delete(fdb.db.Foods, id)
	return nil
}

func (fdb *foodRepo) UpdateFood(new *models.Food) error {
	fdb.db.FoodMu.Lock()
	defer fdb.db.FoodMu.Unlock()
	if _, find := fdb.db.Foods[new.Id]; !find {
		return ErrorInvalidFID
	}

	fdb.db.Foods[new.Id] = new
	return nil
}
