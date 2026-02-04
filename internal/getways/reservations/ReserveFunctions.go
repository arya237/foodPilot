package reservations

import "time"

type ReserveFunctions interface {
	GetProperSelfID(token string) (map[string]int, error)
	GetAccessToken(studentNumber string, password string) (string, error)
	GetFoodProgram(token string, selfID int, startDate time.Time) (*WeekFood, error)
	ReserveFood(token string, meal ReserveModel) (string, error)
}
