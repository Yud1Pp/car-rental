package model

type Customer struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(100);not null"`
	NIK         string `json:"nik" gorm:"type:varchar(20);unique;not null"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20);not null"`
}