package config

import (
	"os"

	"github.com/rs/cors"
)

// GetPort .
func GetPort() string {
	if os.Getenv("PORT") == "" {
		return "8000" //localhost
	}
	return os.Getenv("PORT")
}

// GetCors to config server.
func GetCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "*", "http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "X-Requested-With", "Authorization"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
}
