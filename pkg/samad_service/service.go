package pkg

import (
	"fmt"
	"strconv"

	"github.com/arya237/foodPilot/pkg"
)

func SeperateLunchsDinners(ProgramWeekFoodList []interface{}) (pkg.WeekFood) {

	Week := pkg.WeekFood{
		DailyFood: make(map[string]map[string][]pkg.ReserveModel), 
	}

	for i := range ProgramWeekFoodList {

		dailyFoodProgram := ProgramWeekFoodList[i]
		dailyFoodProgramList, _ := dailyFoodProgram.([]interface{})
		
		for j := range dailyFoodProgramList {

			

			new, _ := dailyFoodProgramList[j].(map[string]interface{})
			
			
			meal := dailyFoodProgramList[j].(map[string]interface{})
			
			if meal["mealTypeId"] == float64(1) {
				
				var mealInfo pkg.ReserveModel
				
				mealInfo.FoodName = meal["foodName"].(string)
				mealInfo.ProgramId = strconv.FormatFloat(meal["programId"].(float64), 'f', -1, 64)
				mealInfo.FoodTypeId = strconv.FormatFloat(meal["foodTypeId"].(float64), 'f', -1, 64)
				mealInfo.MealTypeId = strconv.FormatFloat(meal["mealTypeId"].(float64), 'f', -1, 64)
				
				if _, ok := Week.DailyFood[new["dayTranslated"].(string)]; !ok{

					Week.DailyFood[new["dayTranslated"].(string)] = make(map[string][]pkg.ReserveModel)
				}

				Week.DailyFood[new["dayTranslated"].(string)]["lunch"] = append(Week.DailyFood[new["dayTranslated"].(string)]["lunch"], mealInfo)
				fmt.Println("salam")
			}

			if meal["mealTypeId"] == float64(5) {

				var mealInfo pkg.ReserveModel

				mealInfo.FoodName = meal["foodName"].(string)
				mealInfo.ProgramId = strconv.FormatFloat(meal["programId"].(float64), 'f', -1, 64)
				mealInfo.FoodTypeId = strconv.FormatFloat(meal["foodTypeId"].(float64), 'f', -1, 64)
				mealInfo.MealTypeId = strconv.FormatFloat(meal["mealTypeId"].(float64), 'f', -1, 64)
				
				if _, ok := Week.DailyFood[new["dayTranslated"].(string)]; !ok{

					Week.DailyFood[new["dayTranslated"].(string)] = make(map[string][]pkg.ReserveModel)
				}
				
				Week.DailyFood[new["dayTranslated"].(string)]["dinner"] = append(Week.DailyFood[new["dayTranslated"].(string)]["dinner"], mealInfo)
			}
		}
	}

	return Week
}