package main

import (
	"log"
	"os"

	"ecommerce-microservices/db"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func main() {
	db.SetupDB()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the env file ")
	}
	port := os.Getenv("DB_PORT")
	app := fiber.New()
	app.Listen(":" + port)

}
