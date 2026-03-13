package service

import (
	"github.com/Yud1Pp/car-rental/internal/model"
	"github.com/Yud1Pp/car-rental/internal/repository"
)

type CarService interface {
	GetAll() ([]model.Car, error)
	GetByID(id uint) (*model.Car, error)
	Create(car *model.Car) error
	Update(car *model.Car) error
	Delete(id uint) error
}

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo: repo}
}

func (s *carService) GetAll() ([]model.Car, error) {
	return s.repo.GetAll()
}

func (s *carService) GetByID(id uint) (*model.Car, error) {
	return s.repo.GetByID(id)
}

func (s *carService) Create(car *model.Car) error {
	return s.repo.Create(car)
}

func (s *carService) Update(car *model.Car) error {
	return s.repo.Update(car)
}

func (s *carService) Delete(id uint) error {
	return s.repo.Delete(id)
}