package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerGetComicByID(w http.ResponseWriter, r *http.Request) {
	// Extract comic ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/comics/")
	if path == "" || path == "/" {
		respondWithError(w, http.StatusBadRequest, "Missing comic ID")
		return
	}

	comicID := path

	// Get comic by ID using sqlc-generated function
	comic, err := cfg.DB.GetComicByID(context.Background(), comicID)
	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Comic not found")
		return
	} else if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Convert comic to JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(comic); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

}
