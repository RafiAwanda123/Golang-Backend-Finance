package main

import (
	"fmt"
	"log"

	"github.com/RafiAwanda123/Finance-UMKM/config"
	"github.com/RafiAwanda123/Finance-UMKM/database"
	"github.com/RafiAwanda123/Finance-UMKM/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// Initialize database
	database.InitDB(cfg)
	defer database.DB.Close()

	// Setup router
	router := routes.SetupRouter()

	// Start server (AMAN untuk development)
	port := cfg.ServerPort
	if port == 0 {
		port = 8080 // Default port jika tidak ada di .env
	}

	fmt.Printf("Server berjalan di http://localhost:%d\n", port)
	router.Run(fmt.Sprintf("localhost:%d", port)) // <-- Revisi di sini
}
