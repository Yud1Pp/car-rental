package handler

import (
	"errors"
	"strconv"

	"github.com/Yud1Pp/car-rental/internal/model"
	"github.com/Yud1Pp/car-rental/internal/service"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type BookingHandler struct {
	service service.BookingService
}

func NewBookingHandler(service service.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) handleDatabaseError(c fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "booking not found",
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

// GetBookings godoc
// @Summary Get all bookings
// @Description Retrieve list of all bookings
// @Tags bookings
// @Produce json
// @Success 200 {array} model.Booking
// @Router /bookings [get]
func (h *BookingHandler) GetBookings(c fiber.Ctx) error {

	bookings, err := h.service.GetAll()
	if err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.JSON(bookings)
}

// GetBookingByID godoc
// @Summary Get booking by ID
// @Description Retrieve booking details by ID
// @Tags bookings
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} model.Booking
// @Failure 404 {object} map[string]string
// @Router /bookings/{id} [get]
func (h *BookingHandler) GetBookingByID(c fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	booking, err := h.service.GetByID(uint(id))
	if err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.JSON(booking)
}

// CreateBooking godoc
// @Summary Create new booking
// @Description Create a new booking record
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body model.BookingRequest true "Booking Data"
// @Success 201 {object} model.Booking
// @Failure 400 {object} map[string]string
// @Router /bookings [post]
func (h *BookingHandler) CreateBooking(c fiber.Ctx) error {

	var booking model.Booking

	if err := c.Bind().Body(&booking); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.service.Create(&booking); err != nil {
		return h.handleDatabaseError(c, err)
	}

	createdBooking, err := h.service.GetByID(booking.ID)
	if err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.Status(201).JSON(createdBooking)
}

// UpdateBooking godoc
// @Summary Update booking
// @Description Update an existing booking
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Param booking body model.UpdateBookingRequest true "Booking Data"
// @Success 200 {object} model.Booking
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /bookings/{id} [put]
func (h *BookingHandler) UpdateBooking(c fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var booking model.Booking

	if err := c.Bind().Body(&booking); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	booking.ID = uint(id)

	if err := h.service.Update(&booking); err != nil {
		return h.handleDatabaseError(c, err)
	}

	updatedBooking, err := h.service.GetByID(booking.ID)
	if err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.JSON(updatedBooking)
}

// DeleteBooking godoc
// @Summary Delete booking
// @Description Delete booking by ID
// @Tags bookings
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /bookings/{id} [delete]
func (h *BookingHandler) DeleteBooking(c fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "booking deleted",
	})
}
