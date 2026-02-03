package postgres

import (
	"database/sql"
	"log"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
)

type identities struct {
	db     *sql.DB
}

func NewIdentities(db *sql.DB) repositories.Identities{
	if db == nil {
		log.Fatal("DB is nil")
	}
	return &identities{
		db: db,
	}
}

func (i *identities) Save(new *models.Identities) (*models.Identities, error) {
	query := `
    INSERT INTO identities (user_id, provider, identifier)
    VALUES ($1, $2, $3)
    RETURNING id`

	err := i.db.QueryRow(query, 
        new.UserID,
        new.Provider,
        new.Identifier,
    ).Scan(&new.ID)

	if err != nil {
        return nil, err
    }
	return  new, nil
}
func (i *identities) GetByProvide(provide models.IdProvider, identifier string) (*models.Identities, error) {
	query := `
	SELECT id, user_id, provider, identifier
	FROM identities
	WHERE provider = $1 and identifier = $2`

	var identity models.Identities
	err := i.db.QueryRow(query, provide, identifier).Scan(
		&identity.ID,
		&identity.UserID,
		&identity.Provider,
		&identity.Identifier,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repositories.ErrorNotFound
		}
		return nil, err
	}
	return &identity, nil
}

