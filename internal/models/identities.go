package models

type IdProvider string

const (
	TELEGRAM IdProvider = "telegram"
)

type Identities struct {
	ID         int        `json:"id"`
	UserID     int        `json:"user_id"`
	Provider   IdProvider `json:"provider"`
	Identifier string     `json:"identifier"`
}
