package model

type Car struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"type:varchar(100);not null"`
	Stock     int    `json:"stock" gorm:"not null"`
	DailyRent int    `json:"daily_rent" gorm:"not null"`
}