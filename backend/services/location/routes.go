package location

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanmeij/go-reservation/services/auth"
	"github.com/jonathanmeij/go-reservation/types"
	"github.com/jonathanmeij/go-reservation/utils"
)

type Handler struct {
	store    types.LocationStore
	useStore types.UserStore
}

func NewHandler(store types.LocationStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, useStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/locations/{id}", h.handleGetEventByID).Methods(http.MethodGet)
	router.HandleFunc("/locations", h.handleGetEvents).Methods(http.MethodGet)

	router.HandleFunc("/locations", auth.WithJWTAuthRole(h.handleCreateEvent, h.useStore, "admin")).Methods(http.MethodPost)
	router.HandleFunc("/locations/{id}", auth.WithJWTAuthRole(h.handleDeleteEvent, h.useStore, "admin")).Methods(http.MethodDelete)
}

func (h *Handler) handleGetEventByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	locationEntity, err := h.store.GetLocationByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	location := locationEntity.ToLocation()

	utils.WriteJSON(w, http.StatusOK, location)
}

func (h *Handler) handleGetEvents(w http.ResponseWriter, r *http.Request) {
	locationEntities, err := h.store.GetLocations()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	locations := make([]types.Location, 0, len(locationEntities))
	for _, locationEntity := range locationEntities {
		locations = append(locations, locationEntity.ToLocation())
	}

	utils.WriteJSON(w, http.StatusOK, locations)
}

func (h *Handler) handleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteLocation(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	locationReq := new(types.CreateLocationRequest)
	if err := utils.ParseJSON(r, locationReq); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	locationEntity := locationReq.ToLocationEntity()

	err := h.store.CreateLocation(locationEntity)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, locationEntity)
}
