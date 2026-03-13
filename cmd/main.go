package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Yud1Pp/car-rental/config"
	_ "github.com/Yud1Pp/car-rental/docs"
	"github.com/Yud1Pp/car-rental/internal/router"
	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

// @title Car Rental API
// @version 1.0
// @description Simple Car Rental API
// @host localhost:3000
// @BasePath /api/v1

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Loading dotenv error")
	}

	config.ConnectDatabase()

	app := fiber.New(fiber.Config{})

	router.SetupRoutes(app)

	app.Get("/swagger/*", swaggo.HandlerDefault)

	port := os.Getenv("APP_PORT")
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
