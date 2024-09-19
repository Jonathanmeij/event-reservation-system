package main

import (
	"fmt"
	"log"

	"github.com/jonathanmeij/go-reservation/cmd/api"
	"github.com/jonathanmeij/go-reservation/configs"
	"github.com/jonathanmeij/go-reservation/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("user=postgres dbname=postgres password=gobank sslmode=disable port=5432 host=localhost TimeZone=Europe/Amsterdam"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&types.EventEntity{},
		&types.LocationEntity{},
		&types.PlannedEventEntity{},
		&types.TicketEntity{},
		&types.UserEntity{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database schema: ", err)
	}

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
