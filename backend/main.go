package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// map handlers to endpoints
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/guess", guessHandler)

	// spin up the http server
	port := ":8080"
	fmt.Printf("Starting server on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // using the defaut router
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Accessed login endpoint")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Accessed guess endpoint")
}
