package router

import (
	"github.com/Yud1Pp/car-rental/config"
	"github.com/Yud1Pp/car-rental/internal/handler"
	"github.com/Yud1Pp/car-rental/internal/repository"
	"github.com/Yud1Pp/car-rental/internal/service"
	"github.com/Yud1Pp/car-rental/internal/utils"
	"github.com/gofiber/fiber/v3"
)

func setupCarRoutes(app *fiber.App) {
	carRepo := repository.NewCarRepository(config.DB)
	carService := service.NewCarService(carRepo)
	carHandler := handler.NewCarHandler(carService)

	api := app.Group(utils.APIBaseURL)
	cars := api.Group("/cars")

	cars.Get("/", carHandler.GetCars)
	cars.Get("/:id", carHandler.GetCarByID)
	cars.Post("/", carHandler.CreateCar)
	cars.Put("/:id", carHandler.UpdateCar)
	cars.Delete("/:id", carHandler.DeleteCar)
}