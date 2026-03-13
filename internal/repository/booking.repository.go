package repository

import (
	"github.com/Yud1Pp/car-rental/internal/model"
	"gorm.io/gorm"
)

type BookingRepository interface {
	GetAll() ([]model.Booking, error)
	GetByID(id uint) (*model.Booking, error)
	Create(booking *model.Booking) error
	Update(booking *model.Booking) error
	Delete(id uint) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) GetAll() ([]model.Booking, error) {

	var bookings []model.Booking

	err := r.db.
		Preload("Customer").
		Preload("Car").
		Find(&bookings).Error

	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *bookingRepository) GetByID(id uint) (*model.Booking, error) {

	var booking model.Booking

	err := r.db.
		Preload("Customer").
		Preload("Car").
		First(&booking, id).Error

	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (r *bookingRepository) Create(booking *model.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) Update(booking *model.Booking) error {

	var existingBooking model.Booking

	if err := r.db.First(&existingBooking, booking.ID).Error; err != nil {
		return err
	}

	result := r.db.Model(&existingBooking).Updates(map[string]any{
		"customer_id": booking.CustomerID,
		"car_id":      booking.CarID,
		"start_rent":  booking.StartRent,
		"end_rent":    booking.EndRent,
		"total_cost":  booking.TotalCost,
		"finished":    booking.Finished,
	})

	return result.Error
}

func (r *bookingRepository) Delete(id uint) error {

	result := r.db.Delete(&model.Booking{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
