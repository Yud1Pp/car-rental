package repository

import (
	"github.com/Yud1Pp/car-rental/internal/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetAll() ([]model.Customer, error)
	GetByID(id uint) (*model.Customer, error)
  Create(customer *model.Customer) error
  Update(customer *model.Customer) error
	Delete(id uint) error
}

type customerRepository struct {
  db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) GetAll() ([]model.Customer, error) {

	var customers []model.Customer

	err := r.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepository) GetByID(id uint) (*model.Customer, error) {

	var customer model.Customer

	err := r.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepository) Create(customer *model.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) Update(customer *model.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) Delete(id uint) error {
	return r.db.Delete(&model.Customer{}, id).Error
}