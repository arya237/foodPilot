package reservations

type Weekday int8
type Meal int8

const (
	Lunch Meal = iota + 1
	Dinner
	Breakfast
)

const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
