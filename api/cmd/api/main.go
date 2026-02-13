package main

import (
	"log"
	"os"

	httptransport "github.com/idoceb00/elogap-api/internal/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // use 8081 to avoid clashing with Java on 8080
	}

	r := httptransport.NewRouter()

	log.Printf("Go API listening on :%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
