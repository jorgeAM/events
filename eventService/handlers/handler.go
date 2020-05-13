package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgeAM/events/eventService/db"
	"github.com/jorgeAM/events/eventService/models"
	"github.com/jorgeAM/events/msgbroker"
	"github.com/jorgeAM/events/msgbroker/events"
)

type eventServiceHandler struct {
	dbHandler db.DatabaseHandler
	emitter   msgbroker.EventEmitter
}

// NewEventHandler create new event handler
func NewEventHandler(dbHandler db.DatabaseHandler, emitter msgbroker.EventEmitter) *eventServiceHandler {
	return &eventServiceHandler{
		dbHandler: dbHandler,
		emitter:   emitter,
	}
}

func (h *eventServiceHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Algo salio mal al leer el body de la petición"))
		return
	}

	bytes, err := h.dbHandler.CreateEvent(&event)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Algo salío mal al tratar de crear evento"))
		return
	}

	ev := events.EventCreated{
		ID:    event.ID.Hex(),
		Name:  event.Name,
		Start: event.StartAt,
	}

	h.emitter.Emit(&ev)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}

func (h *eventServiceHandler) ListEventHandler(w http.ResponseWriter, r *http.Request) {
	events, err := h.dbHandler.ListAvailableEvents()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Algo salío mal al tratar de conseguir los eventos: "))
		return
	}

	bytes, err := json.Marshal(&events)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Algo salío mal al tratar de parsear los eventos"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *eventServiceHandler) GetEventHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	event, err := h.dbHandler.GetEvent(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No pudimos encontrar el evento"))
		return
	}

	bytes, err := json.Marshal(event)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Algo salío mal al tratar de parsear el evento"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
