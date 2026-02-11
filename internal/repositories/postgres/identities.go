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

func (i *identities) ListByProvider(provider models.IdProvider, page, pageSize int) ([]*models.Identities, error) {
	//TODO
	if !provider.IsValid() {
		return nil, repositories.ErrorBadArgument
	}
	if page <= 0 || pageSize <= 0 {
		return nil, repositories.ErrorBadArgument
	}

	var identities []*models.Identities
	offset := (page - 1) * pageSize
	
	query := `
	SELECT id, user_id, provider, identifier
	FROM identities
	WHERE provider = $1
	ORDER BY id
	LIMIT $2 OFFSET $3`

	rows, err := i.db.Query(query, provider, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var identity models.Identities
		err := rows.Scan(
			&identity.ID,
			&identity.UserID,
			&identity.Provider,
			&identity.Identifier,
		)
		if err != nil {
			return nil, err
		}
		identities = append(identities, &identity)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return identities, nil
}

