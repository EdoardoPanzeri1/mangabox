package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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

func (cfg *apiConfig) handlerVolumeDetail(w http.ResponseWriter, r *http.Request) {
	volumeID := r.URL.Query().Get("id")
	if volumeID == "" {
		respondWithError(w, http.StatusBadRequest, "Missing volume ID")
		return
	}

	apiURL := fmt.Sprintf("%s/volume/4050-%s/?api_key=%s&format=json", cfg.BaseURL, volumeID, cfg.APIKey)

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch from ComicVine API")
		return
	}
	defer resp.Body.Close()

	var result struct {
		Results Manga `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to decode response from ComicVine API")
		return
	}

	respondWithJSON(w, http.StatusOK, result.Results)
}
