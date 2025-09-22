package main

import (

	"github.com/arya237/foodPilot/pkg/logger"
)

func main(){
	log := logger.New("main logger")
	log.Info("I am alive", logger.Field{
		Key: "key",
		Value: "Value",
	})
}