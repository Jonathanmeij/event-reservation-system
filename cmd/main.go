package main

import (
	"log"

	"github.com/jonathanmeij/go-reservation/configs"
	"github.com/jonathanmeij/go-reservation/db"
)

func main() {
	db, err := db.NewDatabase(configs.Envs.DBUser, configs.Envs.DBName, configs.Envs.DBPassword)
	if err != nil {
		log.Fatal(err)
	}

}
