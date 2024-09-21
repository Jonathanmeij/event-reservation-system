package main

import (
	"fmt"
	"log"

	"github.com/jonathanmeij/go-reservation/cmd/api"
	"github.com/jonathanmeij/go-reservation/configs"
	"github.com/jonathanmeij/go-reservation/db"
	"github.com/jonathanmeij/go-reservation/types"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
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
