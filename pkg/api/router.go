package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanmeij/go-reservation/pkg/models"
	"github.com/jonathanmeij/go-reservation/pkg/service"
)

type Server struct {
	listenAddr     string
	accountService *service.AccountService
}

func NewServer(listenAddr string, accountService *service.AccountService) *Server {
	return &Server{
		listenAddr:     listenAddr,
		accountService: accountService,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleCreateAccount)).Methods("POST")

	log.Println("JSON API server listening on", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *Server) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountRequest := new(models.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
		return err
	}

	account := models.Account{
		FirstName: createAccountRequest.FirstName,
		LastName:  createAccountRequest.LastName}

	err := s.accountService.CreateAccount(&account)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, &account)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
		}
	}
}
