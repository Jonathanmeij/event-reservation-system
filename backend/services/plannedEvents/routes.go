package plannedEvents

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanmeij/go-reservation/services/auth"
	"github.com/jonathanmeij/go-reservation/types"
	"github.com/jonathanmeij/go-reservation/utils"
)

type Handler struct {
	store    types.PlannedEventStore
	useStore types.UserStore
}

func NewHandler(store types.PlannedEventStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, useStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/plannedevents/{id}", h.handleGetPlannedEventByID).Methods(http.MethodGet)
	router.HandleFunc("/plannedevents", h.handleGetPlannedEvents).Methods(http.MethodGet)

	router.HandleFunc("/plannedevents", auth.WithJWTAuthRole(h.handleCreatePlannedEvent, h.useStore, "admin")).Methods(http.MethodPost)
}

func (h *Handler) handleGetPlannedEventByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	plannedEventEntity, err := h.store.GetPlannedEventByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	plannedEvent := plannedEventEntity.ToPlannedEvent()

	utils.WriteJSON(w, http.StatusOK, plannedEvent)
}

func (h *Handler) handleGetPlannedEvents(w http.ResponseWriter, r *http.Request) {
	plannedEventEntities, err := h.store.GetPlannedEvents()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	plannedEvents := make([]types.PlannedEvent, 0, len(plannedEventEntities))
	for _, plannedEventEntity := range plannedEventEntities {
		plannedEvents = append(plannedEvents, plannedEventEntity.ToPlannedEvent())
	}

	utils.WriteJSON(w, http.StatusOK, plannedEvents)
}

func (h *Handler) handleCreatePlannedEvent(w http.ResponseWriter, r *http.Request) {
	plannedEventReq := new(types.CreatePlannedEventRequest)
	if err := utils.ParseJSON(r, plannedEventReq); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	plannedEventEntity := plannedEventReq.ToPlannedEvent()

	err := h.store.CreatePlannedEvent(plannedEventEntity)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, plannedEventEntity)
}
