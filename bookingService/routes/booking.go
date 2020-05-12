package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitBookingRoutes init booking routes
func InitBookingRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/events/{id}/booking", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("booking"))
	}).Methods("POST")

	return r
}
