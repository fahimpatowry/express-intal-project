package main

import (
	"log"
	"os"

	"initial_project/internal/config"
	"initial_project/internal/database"
	"initial_project/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration (e.g., PORT)
	cfg := config.LoadConfig()

	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DB")
	if err := database.ConnectMongo(mongoURI); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Insert a sample user
	database.InsertSampleUser(mongoDB)

	// Initialize server
	s := server.NewServer()

	// Add routes
	s.Routes()

	// Start server
	log.Printf("Server running on port %s\n", cfg.Port)
	if err := s.Start(cfg.Port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
