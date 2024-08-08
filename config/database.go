package config

import (
	"haircut-api-go-gin-gorm-postgres/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dsn := "balinuxsv1:123B@l!nuxV1@tcp(127.0.0.1:3306)/haircut?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:1234Usd!@tcp(127.0.0.1:3306)/haircut?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "balinuxsv1:123B@l!nuxV1@tcp(172.16.211.109:3306)/haircut?charset=utf8mb4&parseTime=True&loc=Local"

	// postgres
	dsn := "host=100.72.100.119 user=balinux password=ninja dbname=haircut port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = database.AutoMigrate(&models.Haircut{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
	DB = database
}
