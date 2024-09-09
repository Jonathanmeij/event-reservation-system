package main

import (
	"log"

	"github.com/jonathanmeij/go-reservation/pkg/api"
	"github.com/jonathanmeij/go-reservation/pkg/service"
	"github.com/jonathanmeij/go-reservation/pkg/storage"
)

func main() {
	db, err := storage.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	accountStorage := storage.NewPostGresAccountStorage(db)

	if err := accountStorage.Init(); err != nil {
		log.Fatal(err)
	}

	accountService := service.NewAccountService(accountStorage)

	router := api.NewServer("localhost:3000", accountService)
	router.Run()
}
