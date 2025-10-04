package models

type User struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"-"`
	AutoSave bool     `json:"auto_save"`
	Role     UserRole `json:"role"`
	Token    string   `json:"token"`
}

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)
