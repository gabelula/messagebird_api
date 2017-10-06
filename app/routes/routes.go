package routes

import (
	"github.com/gabelula/messagebird_api/app/handlers"
	"github.com/gorilla/mux"
)

// API is where the routes are being defined.
func API() *mux.Router {

	// router := http.NewServeMux()

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Index).Methods("GET")
	router.HandleFunc("/v1/message", handlers.MessageSend).Methods("POST")

	return router
}
