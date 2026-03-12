package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Yud1Pp/car-rental/config"
	"github.com/Yud1Pp/car-rental/internal/router"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
    log.Fatal("Loading dotenv error")
  }
  
  config.ConnectDatabase()

  app := fiber.New(fiber.Config{})

  router.SetupRoutes(app)

  port := os.Getenv("APP_PORT")
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}