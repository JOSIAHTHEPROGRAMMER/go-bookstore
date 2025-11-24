package main

import (
	"fmt"
	"log"
	"mybookstore/pkg/config"
	"mybookstore/pkg/models"
	"mybookstore/pkg/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env from the project root
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables only")
	}

	// Get connection string
	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		log.Fatalf("CONNECTION_STRING not set")
	}

	// Connect DB
	config.ConnectDatabase(connectionString)
	fmt.Println("Database connected")

	// config.DB.Exec("TRUNCATE TABLE publications")
	// config.DB.Exec("TRUNCATE TABLE books")
	// config.DB.Exec("TRUNCATE TABLE authors")
	// config.DB.Exec("TRUNCATE TABLE publishers")

	// Migrate database
	err = config.DB.AutoMigrate(&models.Book{}, &models.Author{}, &models.Publisher{}, &models.Publication{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Database migrated")

	// Router
	router := mux.NewRouter()

	// Register ALL routes
	routes.BookRoutes(router)
	routes.AuthorRoutes(router)
	routes.PublisherRoutes(router)
	routes.PublicationRoutes(router)

	// Start server
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
