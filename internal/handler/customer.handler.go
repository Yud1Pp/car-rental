package handler

import (
	"strconv"

	"github.com/Yud1Pp/car-rental/internal/model"
	"github.com/Yud1Pp/car-rental/internal/service"
	"github.com/gofiber/fiber/v3"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) GetCustomers(c fiber.Ctx) error {

	customers, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(customers)
}

func (h *CustomerHandler) GetCustomerByID(c fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	customer, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "customer not found",
		})
	}

	return c.JSON(customer)
}

func (h *CustomerHandler) CreateCustomer(c fiber.Ctx) error {

	var customer model.Customer

	if err := c.Bind().Body(&customer); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.service.Create(&customer); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(customer)
}

func (h *CustomerHandler) UpdateCustomer(c fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var customer model.Customer

	if err := c.Bind().Body(&customer); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	customer.ID = uint(id)

	if err := h.service.Update(&customer); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(customer)
}

func (h *CustomerHandler) DeleteCustomer(c fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "customer deleted",
	})
}