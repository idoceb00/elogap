package main

import (
	"log"
	"os"

	httptransport "github.com/idoceb00/elogap-api/internal/http"
)

func main() {
	port := getPort()

	r := httptransport.NewRouter()

	log.Printf("Starting Elogap API on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("failed to start server: %v", err)
	}
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}

	return "8080"
}
