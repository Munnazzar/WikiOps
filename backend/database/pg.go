package database

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		host, port, user, password, name)
	log.Info(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
