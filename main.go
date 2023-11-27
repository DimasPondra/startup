package main

import (
	"log"
	"startup/app"
	"startup/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := app.InitDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	routes.Init(db)
}
