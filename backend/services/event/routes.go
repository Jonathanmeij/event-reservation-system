package event

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
	store     types.EventStore
	userStore types.UserStore
}

func NewHandler(store types.EventStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/events", h.handleGetEvents).Methods(http.MethodGet)
	router.HandleFunc("/events/{id}", h.handleGetEventByID).Methods(http.MethodGet)
	router.HandleFunc("/events-with-planned", h.handleGetEventsWithPlannedEvents).Methods(http.MethodGet)

	router.HandleFunc("/events", auth.WithJWTAuthRole(h.handleCreateEvent, h.userStore, "admin")).Methods(http.MethodPost)
	router.HandleFunc("/events/{id}", auth.WithJWTAuthRole(h.handleDeleteEvent, h.userStore, "admin")).Methods(http.MethodDelete)
	router.HandleFunc("/events/{id}", auth.WithJWTAuthRole(h.handleUpdateEvent, h.userStore, "admin")).Methods(http.MethodPut)
}

func (h *Handler) handleGetEvents(w http.ResponseWriter, r *http.Request) {
	eventEntities, err := h.store.GetEvents()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	events := make([]types.Event, 0, len(eventEntities))
	for _, eventEntity := range eventEntities {
		events = append(events, eventEntity.ToEvent())
	}

	utils.WriteJSON(w, http.StatusOK, events)
}

func (h *Handler) handleGetEventsWithPlannedEvents(w http.ResponseWriter, r *http.Request) {
	eventEntities, err := h.store.GetEventsWithPlannedEvents()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	events := make([]types.Event, 0, len(eventEntities))
	for _, eventEntity := range eventEntities {
		events = append(events, eventEntity.ToEvent())
	}

	utils.WriteJSON(w, http.StatusOK, events)
}

func (h *Handler) handleGetEventByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	EventEntity, err := h.store.GetEventByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	event := EventEntity.ToEvent()

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
		return
	}

	if err := h.store.DeleteEvent(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) handleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetId(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	updateEventRequest := new(types.UpdateEventRequest)
	if err := utils.ParseJSON(r, &updateEventRequest); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	dbEvent, err := h.store.GetEventByID(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	dbEvent.Title = updateEventRequest.Title
	dbEvent.Description = updateEventRequest.Description
	dbEvent.Images = updateEventRequest.Images

	if err := utils.Validate.Struct(dbEvent); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if err := h.store.UpdateEvent(dbEvent); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, dbEvent)
}
