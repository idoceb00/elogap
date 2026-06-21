package main

import (
	"log"

	"github.com/idoceb00/elogap-api/internal/config"
	httptransport "github.com/idoceb00/elogap-api/internal/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	_, err := gorm.Open(postgres.Open(cfg.DB.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("DATABASE: failed to connect database: %v", err)
	}

	log.Println("DATABASE: connection established")

	r := httptransport.NewRouter()

	log.Printf("SERVER: Starting Elogap API on port %s\n", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("SERVER: failed to start server: %v", err)
	}
}
