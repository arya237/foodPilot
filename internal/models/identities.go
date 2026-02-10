package models

type IdProvider string

const (
	TELEGRAM IdProvider = "telegram"
	BALE     IdProvider = "bale"
)

type Identities struct {
	ID         int        `json:"id"`
	UserID     int        `json:"user_id"`
	Provider   IdProvider `json:"provider"`
	Identifier string     `json:"identifier"`
}

func (i IdProvider) IsValid() bool {
	switch i {
	case TELEGRAM, BALE:
		return true
	}
	return false
}