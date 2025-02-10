package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Print server started message
	log.Printf("Server started in port %s", port)

	// Define a basic handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Todo API is running")
	})

	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
