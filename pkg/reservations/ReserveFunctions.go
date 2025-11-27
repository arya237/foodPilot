package reservations

import "time"

type ReserveFunctions interface {
	GetAccessToken(studentNumber string, password string) (string, error)
	GetFoodProgram(token string, startDate time.Time) (*WeekFood, error)
	ReserveFood(token string, meal ReserveModel) (string, error)
}
