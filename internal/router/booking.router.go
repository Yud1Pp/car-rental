package router

import (
	"github.com/Yud1Pp/car-rental/config"
	"github.com/Yud1Pp/car-rental/internal/handler"
	"github.com/Yud1Pp/car-rental/internal/repository"
	"github.com/Yud1Pp/car-rental/internal/service"
	"github.com/Yud1Pp/car-rental/internal/utils"
	"github.com/gofiber/fiber/v3"
)

func setupBookingRoutes(app *fiber.App) {
	bookingRepo := repository.NewBookingRepository(config.DB)
	carRepo := repository.NewCarRepository(config.DB)

	bookingService := service.NewBookingService(config.DB, bookingRepo, carRepo)
	bookingHandler := handler.NewBookingHandler(bookingService)

	api := app.Group(utils.APIBaseURL)
	bookings := api.Group("/bookings")

	bookings.Get("/", bookingHandler.GetBookings)
	bookings.Get("/:id", bookingHandler.GetBookingByID)
	bookings.Post("/", bookingHandler.CreateBooking)
	bookings.Put("/:id", bookingHandler.UpdateBooking)
	bookings.Delete("/:id", bookingHandler.DeleteBooking)
}
