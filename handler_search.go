package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func (cfg *apiConfig) handlerSearchComic(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("q")
	if searchQuery == "" {
		respondWithError(w, http.StatusBadRequest, "Missing search query")
		return
	}

	apiURL := fmt.Sprintf("https://api.jikan.moe/v4/manga?q=%s", url.QueryEscape(searchQuery))

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Error fetching from Jikan API: %v\n", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch from Jikan API")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to read response from Jikan API")
		return
	}

	log.Printf("Raw API response: %s\n", body)

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Error decoding response from Jikan API: %v\n", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to decode responde from Jikan API")
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}
