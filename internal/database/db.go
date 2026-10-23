package database

import (
	"fmt"
	"log"

	"gym-backend/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(host, user, password, dbname, port string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// Исправлено с models на domain
	err = db.AutoMigrate(&domain.User{}, &domain.Equipment{})
	if err != nil {
		log.Fatalf("Ошибка выполнения миграций: %v", err)
	}

	DB = db
	fmt.Println("База данных успешно подключена и смигрирована!")
	return db
}