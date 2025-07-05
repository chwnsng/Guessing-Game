package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chwnsng/Guessing-Game/backend/handlers"
	"github.com/chwnsng/Guessing-Game/backend/middleware"
	"github.com/rs/cors"
)

func main() {
	// Register Handlers to ServeMux
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handlers.LoginHandler)
	// wrap the guess handler in auth middleware
	mux.HandleFunc("/guess", middleware.AuthMiddleware(handlers.GuessHandler))

	// configure CORS middelware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // allow requests from frontend
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	// wrap handler in cors middleware
	handler := c.Handler(mux)

	// spin up the http server
	port := ":8080"
	fmt.Printf("Starting server on port%v\n", port)
	log.Fatal(http.ListenAndServe(port, handler)) // using the defaut router
}
