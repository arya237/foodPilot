package postgres

import (
	"database/sql"
	"errors"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
)

type rateRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewRateRepo(db *sql.DB) repositories.Rate {
	return &rateRepository{
		logger: logger.New("postgress_rate_repo"),
		db:     db,
	}
}

func (r *rateRepository) Save(userID, foodID, score int) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	var Fid, Uid int

	userQuery := ` select id from users where id = $1; `

	err := r.db.QueryRow(userQuery, userID).Scan(&Uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found in database")
		} else {
			r.logger.Info(err.Error())
			return errors.New("cannot get user")
		}
	}

	foodQuery := ` select id from foods where id = $1`

	err = r.db.QueryRow(foodQuery, foodID).Scan(&Fid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("food not found in database")
		} else {
			r.logger.Info(err.Error())
			return errors.New("cannot get food")
		}
	}

	rateQuery := ` INSERT INTO rates (user_id, food_id, score)
 				  VALUES ($1, $2, $3)
 				  ON CONFLICT (user_id, food_id)
 				  DO UPDATE SET score = EXCLUDED.score;
	`

	_, err = r.db.Exec(rateQuery, userID, foodID, Fid)
	if err != nil {
		r.logger.Info(err.Error())
		return errors.New("cannot save rate")
	}

	return nil
}

func (r *rateRepository) GetByUser(userID int) ([]*models.Rate, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	var id int
	checkUserQuery := `select user_id from rates where user_id = $1`
	err := r.db.QueryRow(checkUserQuery).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

		}
	}

	query := ` SELECT user_id, food_id, score
			  FROM rates
 			  WHERE user_id = $1
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		r.logger.Info(err.Error())
		return nil, errors.New("cannot get rates")
	}

	var rates []*models.Rate

	for rows.Next() {
		var rate models.Rate
		err = rows.Scan(&rate.UserID, &rate.FoodID, &rate.Score)
		if err != nil {
			r.logger.Info(err.Error())
			return nil, errors.New("cannot get the user rate")
		}

		rates = append(rates, &rate)
	}

	return rates, nil
}

func (r *rateRepository) Delete(userID, foodID int) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	query := `DELETE FROM rates WHERE user_id = $1 AND food_id = $2`

	_, err := r.db.Exec(query, userID, foodID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("this rate does not exist in database")
		} else {
			r.logger.Info(err.Error())
			return errors.New("cannot delete the rate")
		}
	}

	return nil
}

func (r *rateRepository) Update(userID int, new *models.Rate) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	var fid, uid int
	if err := r.db.QueryRow("select user_id, food_id from rates where user_id = $1 and food_id = $2", userID, new.FoodID).Scan(&uid, &fid); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("rate not found in database")
		} else {
			r.logger.Info(err.Error())
			return errors.New("cannot get the rate")
		}
	}
	query := `UPDATE rates SET score = $3 WHERE user_id = $1 and food_id = $2`
	_, err := r.db.Exec(query, new.Score, uid, fid)
	if err != nil {
		r.logger.Info(err.Error())
		return errors.New("cannot update the rate")
	}
	return nil
}
