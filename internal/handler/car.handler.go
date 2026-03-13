package handler

import (
	"errors"
	"strconv"

	"github.com/Yud1Pp/car-rental/internal/model"
	"github.com/Yud1Pp/car-rental/internal/service"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type CarHandler struct {
	service service.CarService
}

func NewCarHandler(service service.CarService) *CarHandler {
	return &CarHandler{service: service}
}

func (h *CarHandler) handleDatabaseError(c fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "car not found",
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

// GetCars godoc
// @Summary Get all cars
// @Description Retrieve list of all cars
// @Tags cars
// @Produce json
// @Success 200 {array} model.Car
// @Router /cars [get]
func (h *CarHandler) GetCars(c fiber.Ctx) error {

	cars, err := h.service.GetAll()
	if err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.JSON(cars)
}

// GetCarByID godoc
// @Summary Get car by ID
// @Description Retrieve a car by ID
// @Tags cars
// @Produce json
// @Param id path int true "Car ID"
// @Success 200 {object} model.Car
// @Failure 404 {object} map[string]string
// @Router /cars/{id} [get]
func (h *CarHandler) GetCarByID(c fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	car, err := h.service.GetByID(uint(id))
	if err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.JSON(car)
}

// CreateCar godoc
// @Summary Create new car
// @Description Add a new car to the system
// @Tags cars
// @Accept json
// @Produce json
// @Param car body model.CarRequest true "Car Data"
// @Success 201 {object} model.Car
// @Failure 400 {object} map[string]string
// @Router /cars [post]
func (h *CarHandler) CreateCar(c fiber.Ctx) error {

	var car model.Car

	if err := c.Bind().Body(&car); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.service.Create(&car); err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.Status(201).JSON(car)
}

// UpdateCar godoc
// @Summary Update car
// @Description Update existing car data
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "Car ID"
// @Param car body model.CarRequest true "Car Data"
// @Success 200 {object} model.Car
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /cars/{id} [put]
func (h *CarHandler) UpdateCar(c fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var car model.Car

	if err := c.Bind().Body(&car); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	car.ID = uint(id)

	if err := h.service.Update(&car); err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.JSON(car)
}

// DeleteCar godoc
// @Summary Delete car
// @Description Delete a car by ID
// @Tags cars
// @Produce json
// @Param id path int true "Car ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /cars/{id} [delete]
func (h *CarHandler) DeleteCar(c fiber.Ctx) error {

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
		"message": "car deleted",
	})
}
