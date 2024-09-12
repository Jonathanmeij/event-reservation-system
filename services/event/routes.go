package event

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jonathanmeij/go-reservation/types"
	"github.com/jonathanmeij/go-reservation/utils"
)

type Handler struct {
	store types.EventStore
}

func NewHandler(store types.EventStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/events", h.handleGetEvents).Methods(http.MethodGet)
	router.HandleFunc("/events/{id}", h.handleGetEventByID).Methods(http.MethodGet)
	router.HandleFunc("/events", h.handleCreateEvent).Methods(http.MethodPost)
	router.HandleFunc("/events/{id}", h.handleDeleteEvent).Methods(http.MethodDelete)
	router.HandleFunc("/events", h.handleUpdateEvent).Methods(http.MethodPut)
}

func (h *Handler) handleGetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.store.GetEvents()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, events)
}

func (h *Handler) handleGetEventByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	event, err := h.store.GetEventByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, event)
}

func (h *Handler) handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	createEventRequest := new(types.CreateEventRequest)
	if err := utils.ParseJSON(r, &createEventRequest); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	event := createEventRequest.ToEvent()

	if err := utils.Validate.Struct(event); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if err := h.store.CreateEvent(event); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, event)
}

func (h *Handler) handleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := h.store.DeleteEvent(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) handleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	updateEventRequest := new(types.UpdateEventRequest)
	if err := utils.ParseJSON(r, &updateEventRequest); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	event := updateEventRequest.ToEvent()

	if err := utils.Validate.Struct(event); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if err := h.store.UpdateEvent(event); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, event)
}
