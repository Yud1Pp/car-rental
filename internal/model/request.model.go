package model

import "time"

type CustomerRequest struct {
	Name        string `json:"name" example:"Eko Purwanto"`
	NIK         string `json:"nik" example:"87690234567"`
	PhoneNumber string `json:"phone_number" example:"08769788865"`
}

type CarRequest struct {
	Name      string `json:"name" example:"Toyota Avanza"`
	Stock     int    `json:"stock" example:"5"`
	DailyRent int    `json:"daily_rent" example:"350000"`
}

type BookingRequest struct {
	CustomerID uint      `json:"customer_id" example:"1"`
	CarID      uint      `json:"car_id" example:"2"`
	StartRent  time.Time `json:"start_rent" example:"2026-03-13T09:00:00Z"`
	EndRent    time.Time `json:"end_rent" example:"2026-03-15T09:00:00Z"`
}

type UpdateBookingRequest struct {
	CustomerID uint      `json:"customer_id" example:"1"`
	CarID      uint      `json:"car_id" example:"2"`
	StartRent  time.Time `json:"start_rent" example:"2026-03-13T09:00:00Z"`
	EndRent    time.Time `json:"end_rent" example:"2026-03-15T09:00:00Z"`
	Finished   bool      `json:"finished" example:"false"`
}
