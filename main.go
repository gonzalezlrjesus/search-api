package main

import (
	"log"
	"net/http"

	"github.com/gonzalezlrjesus/search-api/config"
	"github.com/gonzalezlrjesus/search-api/routes"
)

func main() {
	// Routes
	routes := routes.Routes()

	// Server port reading
	port := config.GetPort()

	// CORS configuration
	cors := config.GetCors()

	handler := cors.Handler(routes)

	// Server listening
	err := http.ListenAndServe(":"+port, handler) //Launch the app.
	if err != nil {
		log.Print(err)
	}
}
