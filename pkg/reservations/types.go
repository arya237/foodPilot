package reservations

type Weekday int8
type Meal int8

const (
	Lunch Meal = iota + 1
	Dinner
	Breakfast
)

const (
	Saturday Weekday = iota + 1
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
)

// String returns the canonical name for the Meal value.
func (m Meal) String() string {
	switch m {
	case Lunch:
		return "Lunch"
	case Dinner:
		return "Dinner"
	case Breakfast:
		return "Breakfast"
	default:
		return "UnknownMeal"
	}
}

// String returns the canonical name for the Weekday value.
func (w Weekday) String() string {
	switch w {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	default:
		return "UnknownWeekday"
	}
}
