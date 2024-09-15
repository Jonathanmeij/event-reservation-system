package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanmeij/go-reservation/services/event"
	"github.com/jonathanmeij/go-reservation/services/user"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	eventStore := event.NewStore(s.db)
	eventHandler := event.NewHandler(eventStore, userStore)
	eventHandler.RegisterRoutes(subrouter)

	//static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		Debug:            true,
	})
	handler := c.Handler(router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, handler)
}
