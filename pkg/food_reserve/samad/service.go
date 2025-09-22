package samad

import (
	"strconv"
	food "github.com/arya237/foodPilot/pkg/food_reserve"
)

func CreateWeekFood(ProgramWeekFoodList []interface{}) food.WeekFood {

	Week := food.WeekFood{
		DailyFood: make(map[food.Weekday]map[food.Meal][]food.ReserveModel),
	}

	for i := range ProgramWeekFoodList {

		dailyFoodProgram := ProgramWeekFoodList[i]
		dailyFoodProgramList, _ := dailyFoodProgram.([]interface{})

		for j := range dailyFoodProgramList {
			new, _ := dailyFoodProgramList[j].(map[string]interface{})

			meal := dailyFoodProgramList[j].(map[string]interface{})

			var mealInfo food.ReserveModel

			mealInfo.FoodName = meal["foodName"].(string)
			mealInfo.ProgramId = strconv.FormatFloat(meal["programId"].(float64), 'f', -1, 64)
			mealInfo.FoodTypeId = strconv.FormatFloat(meal["foodTypeId"].(float64), 'f', -1, 64)
			mealInfo.MealTypeId = strconv.FormatFloat(meal["mealTypeId"].(float64), 'f', -1, 64)

			var day food.Weekday
			var m food.Meal
			
			switch new["dayTranslated"].(string) {
				case "Saturday":
					day = food.Saturday
				case "Sunday":
					day = food.Sunday
				case "Monday":
					day = food.Monday
				case "Tuesday":
					day = food.Tuesday
				case "Wednesday":
					day = food.Wednesday
				case "Thursday":
					day = food.Thursday
				case "Friday":
					day = food.Friday
			}

			switch meal["mealTypeId"] {
			case float64(1):
				m = food.Lunch
				if _, ok := Week.DailyFood[day]; !ok {
					Week.DailyFood[day] = make(map[food.Meal][]food.ReserveModel)
				}
				Week.DailyFood[day][m] = append(Week.DailyFood[day][m], mealInfo)
			case float64(5):
				m = food.Dinner
				if _, ok := Week.DailyFood[day]; !ok {
					Week.DailyFood[day] = make(map[food.Meal][]food.ReserveModel)
				}
				Week.DailyFood[day][m] = append(Week.DailyFood[day][m], mealInfo)
			}
		}
	}

	return Week
}
