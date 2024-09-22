package db

import (
	"fmt"
	"log"

	"github.com/jonathanmeij/go-reservation/configs"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable port=%s host=%s TimeZone=Europe/Amsterdam", configs.Envs.DBUser, configs.Envs.DBName, configs.Envs.DBPassword, configs.Envs.DBPort, configs.Envs.DBHost)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
