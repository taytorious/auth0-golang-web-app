package main

import (
	"github.com/taytorious/cavavin/01-Login/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	app.Init()
	StartServer()

}
