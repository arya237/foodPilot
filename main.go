package main

import (
	"log"

	"github.com/arya237/foodPilot/cmd"
)

func main() {

	app, err := cmd.NewApp()
	if err != nil {
		log.Print(err.Error())
	}

	log.Print(app.Run("localhost:8080"))
}
