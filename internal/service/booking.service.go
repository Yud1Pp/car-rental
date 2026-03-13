package service

import (
	"errors"

	"github.com/Yud1Pp/car-rental/internal/model"
	"github.com/Yud1Pp/car-rental/internal/repository"
)

type BookingService interface {
	GetAll() ([]model.Booking, error)
	GetByID(id uint) (*model.Booking, error)
	Create(booking *model.Booking) error
	Update(booking *model.Booking) error
	Delete(id uint) error
}

type bookingService struct {
	repo    repository.BookingRepository
	carRepo repository.CarRepository
}

func NewBookingService(repo repository.BookingRepository, carRepo repository.CarRepository) BookingService {
	return &bookingService{
		repo:    repo,
		carRepo: carRepo,
	}
}

func (s *bookingService) GetAll() ([]model.Booking, error) {
	return s.repo.GetAll()
}

func (s *bookingService) GetByID(id uint) (*model.Booking, error) {
	return s.repo.GetByID(id)
}

func (s *bookingService) Create(booking *model.Booking) error {

	car, err := s.carRepo.GetByID(booking.CarID)
	if err != nil {
		return err
	}

	if car.Stock <= 0 {
		return errors.New("car stock is not available")
	}

	days := int(booking.EndRent.Sub(booking.StartRent).Hours() / 24)

	if days <= 0 {
		days = 1
	}

	totalCost := days * car.DailyRent
	booking.TotalCost = totalCost

	car.Stock -= 1
	if err := s.carRepo.Update(car); err != nil {
		return err
	}

	return s.repo.Create(booking)
}

func (s *bookingService) Update(booking *model.Booking) error {
	return s.repo.Update(booking)
}

func (s *bookingService) Delete(id uint) error {
	return s.repo.Delete(id)
}