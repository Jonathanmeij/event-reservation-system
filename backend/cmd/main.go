package main

import (
	"fmt"
	"log"

	"github.com/jonathanmeij/go-reservation/cmd/api"
	"github.com/jonathanmeij/go-reservation/configs"
	"github.com/jonathanmeij/go-reservation/db"
)

func main() {
	db, err := db.NewDatabase(configs.Envs.DBUser, configs.Envs.DBName, configs.Envs.DBPassword)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
