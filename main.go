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

	// validate := validator.New()
	// validate.RegisterValidation("email_available", func(fl validator.FieldLevel) bool {
	// 	value, kind, nullable := fl.ExtractType(fl.Field())
	// 	fmt.Println("value : " + value.String())
	// 	fmt.Println("kind : " + kind.String())
	// 	fmt.Println(nullable)

	// 	return true
	// })

	routes.Init(db)
}