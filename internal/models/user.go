package models

type User struct {
	Id              int      `json:"id"`
	Username        string   `json:"username"`
	HashPassword    string   `json:"-"`
	Role            UserRole `json:"role"`
	SamadConnection bool     `json:"samad_connection"`
}

type RestaurantCredentials struct {
	Id          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	AccessToken string `json:"access_token"`
	AutoSave    bool   `json:"auto_save"`
}

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)
