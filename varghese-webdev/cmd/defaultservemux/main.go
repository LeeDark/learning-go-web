package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler logic into a Closure
func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}

func messageHandlerClosure(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}

func main() {
	http.HandleFunc("/welcome", messageHandler)
	http.HandleFunc("/message", messageHandlerClosure("net/http is awesome").ServeHTTP)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}