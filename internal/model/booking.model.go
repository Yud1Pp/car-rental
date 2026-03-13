package model

import "time"

type Booking struct {
	ID uint `json:"id" gorm:"primaryKey"`

	CustomerID uint     `json:"customer_id"`
	Customer   Customer `json:"customer" gorm:"foreignKey:CustomerID"`

	CarID uint `json:"car_id"`
	Car   Car  `json:"car" gorm:"foreignKey:CarID"`

	StartRent time.Time `json:"start_rent"`
	EndRent   time.Time `json:"end_rent"`

	TotalCost int  `json:"total_cost"`
	Finished  bool `json:"finished" gorm:"default:false"`
}
