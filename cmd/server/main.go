package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/samezzz/hyperchat/internal/handlers"
)

func main() {
	// Set up the server with the necessary address
	srv := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/send-message", handlers.MessageHandler)
	http.HandleFunc("/status-callback", handlers.StatusCallbackHandler)

	// Start the server in a goroutine to allow graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")

	// Shut down the server gracefully
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
