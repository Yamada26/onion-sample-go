package main

import (
	"log"
	"onion-sample-go/presentation"
)

func main() {
	router := presentation.NewRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
