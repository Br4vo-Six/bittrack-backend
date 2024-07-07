package main

import (
	"log"
	"os"

    "bittrack-backend/controllers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return
	}
	
	controllers.InitConfig(
		os.Getenv("API_PORT"),
	)

	app := controllers.App{}
	app.Initialize()

	app.RunServer()
}
