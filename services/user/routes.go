package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jonathanmeij/go-reservation/services/auth"
	"github.com/jonathanmeij/go-reservation/types"
	"github.com/jonathanmeij/go-reservation/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(userStore types.UserStore) *Handler {
	return &Handler{store: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

}

func (h *Handler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	registerRequest := new(types.RegisterRequest)
	if err := utils.ParseJSON(r, registerRequest); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(registerRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if _, err := h.store.GetUserByEmail(registerRequest.Email); err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("account with email %s already exists", registerRequest.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(registerRequest.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	user := types.NewUser(registerRequest.FirstName, registerRequest.LastName, registerRequest.Email, hashedPassword)
	err = h.store.CreateUser(*user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
