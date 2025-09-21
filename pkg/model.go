package pkg


type ReserveModel struct {
	ProgramId  string
	FoodTypeId string
	MealTypeId string
	FoodName string
}

type WeekFood struct{
	DailyFood map[string]map[string][]ReserveModel
}
