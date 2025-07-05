package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"WorkRESTAPI/internal/server"
)

// Database connection details
var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	schema   = os.Getenv("DB_SCHEMA")
)

// Helper function to get environment variable with default value
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	// Build database URL
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		username, password, host, port, database, schema)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "healthy",
			"service": "WorkRESTAPI",
		})
	})

	// Register all routes from internal
	server.RegisterRoutes(e)

	// Connect to the database
	config, err := pgx.ParseConnectionString(databaseURL)
	if err != nil {
		log.Fatalf("Failed to parse connection string: %v", err)
	}

	db, err := pgx.Connect(config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	// Start server
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "1323"
	}

	log.Printf("Server starting on port %s", serverPort)
	log.Printf("Database URL: postgres://%s:***@%s:%s/%s", username, host, port, database)
	e.Logger.Fatal(e.Start(":" + serverPort))
}
