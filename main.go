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

	baseURL := os.Getenv("JIKAN_BASE_URL")
	if baseURL == "" {
		log.Fatal("COMICWINE_BASE_URL environment is not set")
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
	}

	mux := http.NewServeMux()

	// External Manga Search Endpoints
	mux.HandleFunc("GET /search", apiCfg.handlerSearchManga)
	mux.HandleFunc("POST /details", apiCfg.handlerGetManga)

	// Manga Catalog Endpoints
	//mux.HandleFunc("GET /mangas", apiCfg.handlerRetrieveCatalog)
	//mux.HandleFunc("POST /mangas", apiCfg.handlerAddToCatalog)
	//mux.HandleFunc("PUT /mangas/{id}", apiCfg.handlerStatusManga)
	//mux.HandleFunc("DELETE /mangas/{id}", apiCfg.handlerDeleteManga)

	// User Authentication and Profile Management
	// mux.HandleFunc("POST /register", apiCfg.handlerRegistration)
	// mux.HandleFunc("POST /login", apiCfg.handlerLogin)
	//mux.HandleFunc("GET /profile", apiCfg.handlerProfileInformation)
	//mux.HandleFunc("PUT /profile", apiCfg.handlerUpdateInformation)

	mux.HandleFunc("/v1/healthz", handlerReadiness)
	mux.HandleFunc("/v1/err", handlerErr)

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
