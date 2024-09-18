package main

import (
	"github.com/samezzz/hyperchat/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/send-message", handlers.MessageHandler)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
