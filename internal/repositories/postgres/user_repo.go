package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}


func (r *UserRepository) Save(newUser *models.User) (*models.User, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	query := `
		INSERT INTO users (username, password, auto_save, role, token) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id
	`

	var id int
	err := r.db.QueryRow(
		query,
		newUser.Username,
		newUser.Password,
		newUser.AutoSave,
		newUser.Role,
		newUser.Token,
	).Scan(&id)

	if err != nil {
		if isDuplicateError(err) {
			return nil, repositories.ErrorDuplicate
		}
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	newUser.Id = id
	return newUser, nil
}

func (r *UserRepository) GetById(id int) (*models.User, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	query := `
		SELECT id, username, password, auto_save, role, token 
		FROM users 
		WHERE id = $1
	`

	var user models.User
	err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.AutoSave,
		&user.Role,
		&user.Token,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repositories.ErrorNotFound
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) GetByUserName(username string) (*models.User, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	query := `
		SELECT id, username, password, auto_save, role, token 
		FROM users 
		WHERE username = $1
	`

	var user models.User
	err := r.db.QueryRow(query, username).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.AutoSave,
		&user.Role,
		&user.Token,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repositories.ErrorNotFound
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	query := `
		SELECT id, username, password, auto_save, role, token 
		FROM users 
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.AutoSave,
			&user.Role,
			&user.Token,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users: %w", err)
	}

	return users, nil
}

func (r *UserRepository) Delete(id int) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}

	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return repositories.ErrorNotFound
	}

	return nil
}

func (r *UserRepository) Update(updatedUser *models.User) error {
	if r.db == nil {
		return errors.New("database connection is nil")
	}

	query := `
		UPDATE users 
		SET username = $1, password = $2, auto_save = $3, role = $4, token = $5 
		WHERE id = $6
	`

	result, err := r.db.Exec(
		query,
		updatedUser.Username,
		updatedUser.Password,
		updatedUser.AutoSave,
		updatedUser.Role,
		updatedUser.Token,
		updatedUser.Id,
	)

	if err != nil {
		if isDuplicateError(err) {
			return repositories.ErrorDuplicate
		}
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return repositories.ErrorNotFound
	}

	return nil
}


func isDuplicateError(err error) bool {
	if err == nil {
		return false
	}

	
	errorStr := err.Error()
	return strings.Contains(errorStr, "23505") ||
		strings.Contains(errorStr, "duplicate key") ||
		strings.Contains(errorStr, "already exists")
}


