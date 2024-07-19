package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (cfg *apiConfig) handlerSearchComic(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("q")
	if searchQuery == "" {
		respondWithError(w, http.StatusBadRequest, "Missing search query")
		return
	}

	apiURL := fmt.Sprintf("%s/search/?api_key=%s&format=json&resources=volume&query=%s", cfg.BaseURL, cfg.APIKey, url.QueryEscape(searchQuery))

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch from ComicVine API")
		return
	}
	defer resp.Body.Close()

	var result ComicVineResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to decode responde from ComicVine API")
		return
	}

	respondWithJSON(w, http.StatusOK, result.Results)
}
