package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jonathanmeij/go-reservation/configs"
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
	router.HandleFunc("/account/login", h.HandleLoginUser).Methods(http.MethodPost)
	router.HandleFunc("/account/register", h.HandleRegisterUser).Methods(http.MethodPost)
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

	dbUser, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	jwtToken, err := auth.CreateJWT(dbUser.ID, []byte(configs.Envs.JWTSecret))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.TokenResponse{Token: jwtToken, FirstName: dbUser.FirstName, LastName: dbUser.LastName, Email: dbUser.Email, Role: dbUser.Role})
}

func (h *Handler) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	loginRequest := new(types.LoginRequest)
	if err := utils.ParseJSON(r, loginRequest); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(loginRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	user, err := h.store.GetUserByEmail(loginRequest.Email)
	if err != nil {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("invalid email or password"))
		return
	}

	if !auth.CheckPasswordHash(user.Password, loginRequest.Password) {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("invalid email or password"))
		return
	}

	jwtToken, err := auth.CreateJWT(user.ID, []byte(configs.Envs.JWTSecret))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.TokenResponse{Token: jwtToken, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Role: user.Role})
}
