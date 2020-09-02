package routes

import (
	"github.com/gonzalezlrjesus/search-api/controllers"
	"github.com/gorilla/mux"
)

// Routes contains the endpoint to use.
func Routes() *mux.Router {
	router := mux.NewRouter()

	// Api version.
	s := router.PathPrefix("/v1").Subrouter()

	// Search endpoint.
	s.HandleFunc("/search", controllers.SearchController).Methods("GET")

	return router
}
