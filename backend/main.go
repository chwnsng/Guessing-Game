package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chwnsng/Guessing-Game/backend/handlers"
	"github.com/chwnsng/Guessing-Game/backend/middleware"
)

func main() {
	// map handlers to endpoints
	http.HandleFunc("/login", handlers.LoginHandler)
	// wrap the guess handler in middleware
	http.HandleFunc("/guess", middleware.AuthMiddleware(handlers.GuessHandler))

	// spin up the http server
	port := ":8080"
	fmt.Printf("Starting server on port%v\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // using the defaut router
}

// // pass in Request as pointer to avoid copying the whole struct
// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Accessed login endpoint")
// }

// func guessHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Accessed guess endpoint")
// }
