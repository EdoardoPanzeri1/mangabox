package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/EdoardoPanzeri1/mangabox/internal/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
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
		log.Fatal("JIKAN_BASE_URL environment is not set")
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
	defer db.Close() // Ensure database connection is closed when application shuts down

	router := mux.NewRouter()

	// External Manga Search Endpoints
	router.HandleFunc("/search", apiCfg.handlerSearchManga).Methods("GET")
	router.HandleFunc("/details", apiCfg.handlerGetManga).Methods("GET")

	// Manga Catalog Endpoints
	router.HandleFunc("/mangas", apiCfg.handlerRetrieveCatalog).Methods("GET")
	router.HandleFunc("/mangas", apiCfg.handlerAddToCatalog).Methods("POST")
	router.HandleFunc("/mangas/{id}", apiCfg.handlerStatusManga).Methods("PUT")
	router.HandleFunc("/mangas/{id}", apiCfg.handlerDeleteManga).Methods("DELETE")

	// User Authentication and Profile Management
	router.HandleFunc("/register", apiCfg.handlerRegistration).Methods("POST")
	router.HandleFunc("/login", apiCfg.handlerLogin).Methods("POST")
	router.HandleFunc("/profile", apiCfg.handlerProfileInformation).Methods("GET")
	router.HandleFunc("/profile", apiCfg.handlerUpdateInformation).Methods("PUT")

	router.HandleFunc("/v1/healthz", handlerReadiness)
	router.HandleFunc("/v1/err", handlerErr)

	// Setup CORS to allow requests from the React app
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // React's dev server
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)

	// Set up HTTP server
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	// Start the server
	log.Printf("Serving file on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
