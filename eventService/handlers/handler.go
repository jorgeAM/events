package handlers

import (
	"net/http"

	"github.com/jorgeAM/events/eventService/db"
)

type eventServiceHandler struct {
	dbHandler db.DatabaseHandler
}

// NewEventHandler create new event handler
func NewEventHandler(dbHandler db.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbHandler,
	}
}

func (h *eventServiceHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("crear evento"))
}

func (h *eventServiceHandler) ListEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("listar evento"))
}

func (h *eventServiceHandler) GetEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscar evento"))
}
