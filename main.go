package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func main() {
	db.setupDB()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the env file ")
	}
	port := os.Getenv("DB_PORT")
	app := fiber.New()
	app.Listen(":" + port)
}
