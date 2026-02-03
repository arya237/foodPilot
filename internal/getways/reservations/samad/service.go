package samad

import (
	"github.com/arya237/foodPilot/internal/getways/reservations"
	"sort"
	"strconv"
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
			case float64(3):
				m = reservations.Breakfast
				if _, ok := Week.DailyFood[day]; !ok {
					Week.DailyFood[day] = make(map[reservations.Meal][]reservations.ReserveModel)
				}
				Week.DailyFood[day][m] = append(Week.DailyFood[day][m], mealInfo)
			}
		}
	}

	sortWeekFood(&Week)
	return Week
}

func sortWeekFood(week *reservations.WeekFood) {
	days := make([]reservations.Weekday, 0, len(week.DailyFood))
	for day := range week.DailyFood {
		days = append(days, day)
	}

	sort.Slice(days, func(i, j int) bool {
		return days[i] < days[j]
	})

	for _, day := range days {
		meals := week.DailyFood[day]
		mealKeys := make([]reservations.Meal, 0, len(meals))
		for m := range meals {
			mealKeys = append(mealKeys, m)
		}
		sort.Slice(mealKeys, func(i, j int) bool {
			return mealKeys[i] < mealKeys[j]
		})
	}
}
