package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"message-processor/api"
	"message-processor/db"
	"message-processor/kafka"
)

func main() {
	// Initialize the database and Kafka
	db.InitDB()
	kafka.InitProducer()
	kafka.InitConsumer()

	// Create a new router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/messages", api.CreateMessageHandler).Methods("POST")
	r.HandleFunc("/stats", api.GetMessageStatsHandler).Methods("GET")

	// Serve static files
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

	// Get the port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Log registered routes
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		log.Printf("Registered route: %s", path)
		return nil
	})

	// Start Kafka consumer in a separate goroutine
	go kafka.ConsumeMessages()

	// Start the server
	log.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
