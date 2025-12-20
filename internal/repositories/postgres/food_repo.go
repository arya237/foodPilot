package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
)

type foodRepo struct {
	db     *sql.DB
	logger logger.Logger
}

func NewFoodRepo(db *sql.DB) repositories.Food {
	return &foodRepo{
		db:     db,
		logger: logger.New("PostGres_food_repo"),
	}
}

func (r *foodRepo) Save(name string) (int, error) {
	if r.db == nil {
		return 0, errors.New("database connection is nil")
	}

	foodQuery := ` INSERT INTO foods (name) VALUES ($1); `

	_, err := r.db.Exec(foodQuery, name)
	if err != nil {
		r.logger.Info("Error inserting food into the database")
		return 0, err
	}

	return 0, nil
}

func (r *foodRepo) GetById(id int) (*models.Food, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	query := `select * from foods where id = $1;`
	var food models.Food
	err := r.db.QueryRow(query, id).Scan(&food.Id, &food.Name)
	if err != nil {
		r.logger.Info(err.Error())
		return nil, errors.New(fmt.Sprintf("error getting food by id %d", id))
	}

	return &food, nil
}

func (r *foodRepo) GetAll() ([]*models.Food, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	query := `select * from foods;`
	rows, err := r.db.Query(query)
	if err != nil {
		r.logger.Info(err.Error())
		return nil, errors.New(fmt.Sprintf("error getting foods from database"))
	}

	defer rows.Close()

	var foods []*models.Food
	for rows.Next() {
		var food models.Food
		err := rows.Scan(&food.Id, &food.Name)
		if err != nil {
			r.logger.Info(err.Error())
			return nil, errors.New(fmt.Sprintf("error getting foods from database"))
		}

		foods = append(foods, &food)
	}

	return foods, nil
}

func (r *foodRepo) Delete(id int) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}

	query := `delete from foods where id = $1;`

	_, err := r.db.Exec(query, id)
	if err != nil {
		r.logger.Info(err.Error())
		return errors.New(fmt.Sprintf("error deleting food with id %d from database", id))
	}

	return nil
}

func (r foodRepo) Update(new *models.Food) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}

	query := `update foods set name = $1 where id = $2;`

	_, err := r.db.Exec(query, new.Name, new.Id)
	if err != nil {
		r.logger.Info(err.Error())
		return errors.New(fmt.Sprintf("error updating food with id %d", new.Id))
	}

	return nil
}
