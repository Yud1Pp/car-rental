package config

import (
	"fmt"
	"os"

	"github.com/Yud1Pp/car-rental/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_PORT"),
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
  })
  if err != nil {
    panic("Failed connect to database: " + err.Error())
  }

  DB = db

  err = DB.AutoMigrate(
		&model.Customer{},
	)

  if err != nil {
		panic("Failed migrate database: " + err.Error())
	}

  fmt.Println("Database connected successfully")
}