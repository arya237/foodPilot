package models

type User struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Role     UserRole `json:"role"`
	AutoSave bool     `json:"auto_save"`
}

type RestaurantCredentials struct {
	Id       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Token    string `json:"token"`
}

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)
