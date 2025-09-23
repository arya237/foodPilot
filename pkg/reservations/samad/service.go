package samad

import (
	"strconv"
	"github.com/arya237/foodPilot/pkg/reservations"
)

func CreateWeekFood(ProgramWeekFoodList []interface{}) reservations.WeekFood {

	Week := reservations.WeekFood{
		DailyFood: make(map[reservations.Weekday]map[reservations.Meal][]reservations.ReserveModel),
	}

	for i := range ProgramWeekFoodList {

		dailyFoodProgram := ProgramWeekFoodList[i]
		dailyFoodProgramList, _ := dailyFoodProgram.([]interface{})

		for j := range dailyFoodProgramList {
			new, _ := dailyFoodProgramList[j].(map[string]interface{})

			meal := dailyFoodProgramList[j].(map[string]interface{})

			var mealInfo reservations.ReserveModel

			mealInfo.FoodName = meal["foodName"].(string)
			mealInfo.ProgramId = strconv.FormatFloat(meal["programId"].(float64), 'f', -1, 64)
			mealInfo.FoodTypeId = strconv.FormatFloat(meal["foodTypeId"].(float64), 'f', -1, 64)
			mealInfo.MealTypeId = strconv.FormatFloat(meal["mealTypeId"].(float64), 'f', -1, 64)

			var day reservations.Weekday
			var m reservations.Meal
			
			switch new["dayTranslated"].(string) {
				case "Saturday":
					day = reservations.Saturday
				case "Sunday":
					day = reservations.Sunday
				case "Monday":
					day = reservations.Monday
				case "Tuesday":
					day = reservations.Tuesday
				case "Wednesday":
					day = reservations.Wednesday
				case "Thursday":
					day = reservations.Thursday
				case "Friday":
					day = reservations.Friday
			}

			switch meal["mealTypeId"] {
			case float64(1):
				m = reservations.Lunch
				if _, ok := Week.DailyFood[day]; !ok {
					Week.DailyFood[day] = make(map[reservations.Meal][]reservations.ReserveModel)
				}
				Week.DailyFood[day][m] = append(Week.DailyFood[day][m], mealInfo)
			case float64(5):
				m = reservations.Dinner
				if _, ok := Week.DailyFood[day]; !ok {
					Week.DailyFood[day] = make(map[reservations.Meal][]reservations.ReserveModel)
				}
				Week.DailyFood[day][m] = append(Week.DailyFood[day][m], mealInfo)
			}
		}
	}

	return Week
}
