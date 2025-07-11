package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	//"github.com/pressly/goose/v3"

	internals "WorkRESTAPI/internal"
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
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		username, password, host, port, database)

	log.Printf("DB Config: user=%s, host=%s, port=%s, db=%s, schema=%s",
		username, host, port, database, schema)
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
	// Connect to the database
	db, err := sql.Open("pgx/v5", databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	// Register all routes from internal
	queries := internals.New(db)
	server.RegisterRoutes(e, queries)

	// Start server
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "1323"
	}

	log.Printf("Server starting on port %s", serverPort)
	log.Printf("Database URL: postgres://%s:***@%s:%s/%s", username, host, port, database)
	e.Logger.Fatal(e.Start(":" + serverPort))
}
