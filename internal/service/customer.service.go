package service

import (
	"github.com/Yud1Pp/car-rental/internal/model"
	"github.com/Yud1Pp/car-rental/internal/repository"
)

type CustomerService interface {
	GetAll() ([]model.Customer, error)
	GetByID(id uint) (*model.Customer, error)
	Create(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id uint) error
}

type customerService struct {
  repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) GetAll() ([]model.Customer, error) {
	return s.repo.GetAll()
}

func (s *customerService) GetByID(id uint) (*model.Customer, error) {
	return s.repo.GetByID(id)
}

func (s *customerService) Create(customer *model.Customer) error {
	return s.repo.Create(customer)
}

func (s *customerService) Update(customer *model.Customer) error {
	return s.repo.Update(customer)
}

func (s *customerService) Delete(id uint) error {
	return s.repo.Delete(id)
}