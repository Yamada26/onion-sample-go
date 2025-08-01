package main

import (
	"log"
	"onion-sample-go/infrastructure"
	"onion-sample-go/presentation"
)

func main() {
	infrastructure.InitDB("app.db")

	router := presentation.NewRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
