package main

import (
	"fmt"
	"log"
	"time"
	// "github.com/arya237/foodPilot/pkg"
	samad "github.com/arya237/foodPilot/pkg/samad"
)

func main(){
	var model samad.Samad
	token, err := model.GetAccessToken("40112358043", "arya1383")

	if err != nil{
		log.Println(err)
		return 
	}

	fmt.Println(token)

	timeStr := "2025-09-20 00:00:00"
	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return
	}

	listFood, err := model.GetFoodProgram(token, t)

	if err != nil{
		log.Println(err)
		return 
	}

	fmt.Println(listFood)

	// res, err := model.ReserveFood(token, pkg.ReserveModel{ProgramId: "494338", FoodTypeId: "599", MealTypeId: "1"})
	// if err != nil{
	// 	log.Println(err)
	// 	return 
	// }
	// fmt.Println(res)
}