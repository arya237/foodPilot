package main

import (
	"log"

	"github.com/arya237/foodPilot/internal/handler"
)


func main(){
	server := handler.New()

	if err := server.Run(":8080"); err != nil {
		log.Println(err)
	}
}