package models

type IdProvider string

const (
	TELEGRAM IdProvider = "telegram"
)

type Identities struct {
	ID         int
	UserID     int
	Provider   IdProvider
	Identifier string
}
