package router

import (
	"github.com/Yud1Pp/car-rental/config"
	"github.com/Yud1Pp/car-rental/internal/handler"
	"github.com/Yud1Pp/car-rental/internal/repository"
	"github.com/Yud1Pp/car-rental/internal/service"
	"github.com/Yud1Pp/car-rental/internal/utils"
	"github.com/gofiber/fiber/v3"
)

func setupCustomerRoutes(app *fiber.App) {
	customerRepo := repository.NewCustomerRepository(config.DB)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	api := app.Group(utils.APIBaseURL)
	customers := api.Group("/customers")

	customers.Get("/", customerHandler.GetCustomers)
	customers.Get("/:id", customerHandler.GetCustomerByID)
	customers.Post("/", customerHandler.CreateCustomer)
	customers.Put("/:id", customerHandler.UpdateCustomer)
	customers.Delete("/:id", customerHandler.DeleteCustomer)
}
