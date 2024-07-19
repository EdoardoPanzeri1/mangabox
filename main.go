package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/EdoardoPanzeri1/mangabox/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB      *database.Queries
	BaseURL string
	APIKey  string
}

func main() {
	// Load environment variables from .env file
	godotenv.Load(".env")

	// Get required environment variables
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port environment variable is not set")
	}

	baseURL := os.Getenv("COMICWINE_BASE_URL")
	if baseURL == "" {
		log.Fatal("COMICWINE_BASE_URL environment is not set")
	}

	apiKey := os.Getenv("COMICWINE_API_KEY")
	if apiKey == "" {
		log.Fatal("COMICWINE_API_KEY environment is not set")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment is not set")
	}

	// Initialize database connection
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize sqlc-generated queries
	dbQueries := database.New(db)

	// Set up API configuration
	apiCfg := apiConfig{
		DB:      dbQueries,
		BaseURL: baseURL,
		APIKey:  apiKey,
	}

	mux := http.NewServeMux()

	// Register handler functions
	mux.HandleFunc("GET /comics/", apiCfg.handlerGetComicByID)
	mux.HandleFunc("GET /search/comics", apiCfg.handlerSearchComic)
	mux.HandleFunc("GET /search/comics/", apiCfg.handlerVolumeDetail)

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	// Set up HTTP server
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// Start the server
	log.Printf("Serving file on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
