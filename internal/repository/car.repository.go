package repository

import (
	"github.com/Yud1Pp/car-rental/internal/model"
	"gorm.io/gorm"
)

type CarRepository interface {
	GetAll() ([]model.Car, error)
	GetByID(id uint) (*model.Car, error)
	Create(car *model.Car) error
	Update(car *model.Car) error
	Delete(id uint) error
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db: db}
}

func (r *carRepository) GetAll() ([]model.Car, error) {

	var cars []model.Car

	err := r.db.Find(&cars).Error
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (r *carRepository) GetByID(id uint) (*model.Car, error) {

	var car model.Car

	err := r.db.First(&car, id).Error
	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (r *carRepository) Create(car *model.Car) error {
	return r.db.Create(car).Error
}

func (r *carRepository) Update(car *model.Car) error {

	var existingCar model.Car
	if err := r.db.First(&existingCar, car.ID).Error; err != nil {
		return err
	}

	result := r.db.Model(&existingCar).Updates(map[string]any{
		"name":       car.Name,
		"stock":      car.Stock,
		"daily_rent": car.DailyRent,
	})

	return result.Error
}

func (r *carRepository) Delete(id uint) error {

	result := r.db.Delete(&model.Car{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}