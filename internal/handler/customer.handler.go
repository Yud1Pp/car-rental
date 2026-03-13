package handler

import (
	"errors"
	"strconv"

	"github.com/Yud1Pp/car-rental/internal/model"
	"github.com/Yud1Pp/car-rental/internal/service"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) handleDatabaseError(c fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "customer not found",
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

// GetCustomers godoc
// @Summary Get all customers
// @Description Retrieve a list of all customers
// @Tags customers
// @Produce json
// @Success 200 {array} model.Customer
// @Router /customers [get]
func (h *CustomerHandler) GetCustomers(c fiber.Ctx) error {

	customers, err := h.service.GetAll()
	if err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.JSON(customers)
}

// GetCustomerByID godoc
// @Summary Get customer by ID
// @Description Retrieve a single customer by ID
// @Tags customers
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} model.Customer
// @Failure 404 {object} map[string]string
// @Router /customers/{id} [get]
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
		return h.handleDatabaseError(c, err)
	}

	return c.JSON(customer)
}

// CreateCustomer godoc
// @Summary Create new customer
// @Description Create a new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body model.CustomerRequest true "Customer Data"
// @Success 201 {object} model.Customer
// @Failure 400 {object} map[string]string
// @Router /customers [post]
func (h *CustomerHandler) CreateCustomer(c fiber.Ctx) error {
	var customer model.Customer

	if err := c.Bind().Body(&customer); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.service.Create(&customer); err != nil {
		return h.handleDatabaseError(c, err)
	}

	return c.Status(201).JSON(customer)
}

// UpdateCustomer godoc
// @Summary Update customer
// @Description Update an existing customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param customer body model.CustomerRequest true "Customer Data"
// @Success 200 {object} model.Customer
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /customers/{id} [put]
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
		return h.handleDatabaseError(c, err)
	}

	return c.JSON(customer)
}

// DeleteCustomer godoc
// @Summary Delete customer
// @Description Delete customer by ID
// @Tags customers
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /customers/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(c fiber.Ctx) error {

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
		"message": "customer deleted",
	})
}
