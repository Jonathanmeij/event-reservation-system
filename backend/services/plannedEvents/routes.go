package plannedEvents

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanmeij/go-reservation/types"
	"github.com/jonathanmeij/go-reservation/utils"
)

type Handler struct {
	store types.PlannedEventStore
}

func NewHandler(store types.PlannedEventStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/events/{id}", h.handleGetEventByID).Methods(http.MethodGet)
}

func (h *Handler) handleGetEventByID(w http.ResponseWriter, r *http.Request) {
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

	plannedEvent := types.NewPlannedEvent(plannedEventEntity)

	utils.WriteJSON(w, http.StatusOK, plannedEvent)
}
