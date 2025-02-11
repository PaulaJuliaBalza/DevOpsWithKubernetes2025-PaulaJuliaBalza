package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	// Generate a random UUID on startup
	randomString := uuid.New().String()

	// Output every 5 seconds with a timestamp
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		fmt.Printf("%s: %s\n", t.UTC().Format(time.RFC3339Nano), randomString)
	}
}
