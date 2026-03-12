package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Yud1Pp/car-rental/config"
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

  app.Get("/ping", func(c fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "pong"})
  })

  port := os.Getenv("APP_PORT")
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}