package service

import (
	"errors"
	"time"

	"github.com/Yud1Pp/car-rental/internal/model"
	"github.com/Yud1Pp/car-rental/internal/repository"
	"gorm.io/gorm"
)

type BookingService interface {
	GetAll() ([]model.Booking, error)
	GetByID(id uint) (*model.Booking, error)
	Create(booking *model.Booking) error
	Update(booking *model.Booking) error
	Delete(id uint) error
}

type bookingService struct {
	db      *gorm.DB
	repo    repository.BookingRepository
	carRepo repository.CarRepository
}

func NewBookingService(db *gorm.DB, repo repository.BookingRepository, carRepo repository.CarRepository) BookingService {
	return &bookingService{
		db:      db,
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
	return s.db.Transaction(func(tx *gorm.DB) error {
		txBookingRepo := repository.NewBookingRepository(tx)
		txCarRepo := repository.NewCarRepository(tx)

		car, err := txCarRepo.GetByID(booking.CarID)
		if err != nil {
			return err
		}

		if car.Stock <= 0 {
			return errors.New("car stock is not available")
		}

		booking.TotalCost = calculateBookingTotalCost(booking.StartRent, booking.EndRent, car.DailyRent)

		car.Stock -= 1
		if err := txCarRepo.Update(car); err != nil {
			return err
		}

		return txBookingRepo.Create(booking)
	})
}

func (s *bookingService) Update(booking *model.Booking) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		txBookingRepo := repository.NewBookingRepository(tx)
		txCarRepo := repository.NewCarRepository(tx)

		existingBooking, err := txBookingRepo.GetByID(booking.ID)
		if err != nil {
			return err
		}

		if booking.CustomerID == 0 {
			booking.CustomerID = existingBooking.CustomerID
		}

		if booking.CarID == 0 {
			booking.CarID = existingBooking.CarID
		}

		if booking.StartRent.IsZero() {
			booking.StartRent = existingBooking.StartRent
		}

		if booking.EndRent.IsZero() {
			booking.EndRent = existingBooking.EndRent
		}

		bookingCarChanged := booking.CarID != existingBooking.CarID
		existingReserved := !existingBooking.Finished
		newReserved := !booking.Finished

		if existingReserved && (bookingCarChanged || !newReserved) {
			oldCar, err := txCarRepo.GetByID(existingBooking.CarID)
			if err != nil {
				return err
			}

			oldCar.Stock += 1
			if err := txCarRepo.Update(oldCar); err != nil {
				return err
			}
		}

		currentCar, err := txCarRepo.GetByID(booking.CarID)
		if err != nil {
			return err
		}

		if newReserved && (bookingCarChanged || !existingReserved) {
			if currentCar.Stock <= 0 {
				return errors.New("car stock is not available")
			}

			currentCar.Stock -= 1
			if err := txCarRepo.Update(currentCar); err != nil {
				return err
			}
		}

		booking.TotalCost = calculateBookingTotalCost(booking.StartRent, booking.EndRent, currentCar.DailyRent)

		return txBookingRepo.Update(booking)
	})
}

func (s *bookingService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func calculateBookingTotalCost(startRent, endRent time.Time, dailyRent int) int {
	days := int(endRent.Sub(startRent).Hours() / 24)
	if days <= 0 {
		days = 1
	}

	return days * dailyRent
}
