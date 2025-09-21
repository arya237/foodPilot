package pkg

import "time"

type ReserveModel struct {
	ProgramId  string
	FoodName   string
	FoodTypeId string
	MealTypeId string
}

type RequiredFunctions interface {
	GetAccessToken(studentNumber string, password string) (string, error)
	GetFoodProgram(token string, startDate time.Time) (string, error)
	ReserveFood(token string, meal ReserveModel) (string, error)
}
