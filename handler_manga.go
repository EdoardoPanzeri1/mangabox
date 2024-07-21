package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerGetManga(w http.ResponseWriter, r *http.Request) {
	// Extract comic ID from URL parameters
	mangaIDStr := r.URL.Query().Get("id")
	if mangaIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "id parameter is missing")
		return
	}

	mangaID, err := strconv.Atoi(mangaIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	// Fetch manga details from the external API
	manga, err := fetchMangaDetailsFromAPI(mangaID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "manga not found")
		return
	}

	// Serialize manga details to JSON and write response
	respondWithJSON(w, http.StatusOK, manga)
}

func fetchMangaDetailsFromAPI(mangaID int) (Manga, error) {
	apiURL := fmt.Sprintf("http://api.jikan.moe/v4/manga/%d", mangaID)
	response, err := http.Get(apiURL)
	if err != nil {
		return Manga{}, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return Manga{}, fmt.Errorf("manga not found for ID %d", mangaID)
	}

	if response.StatusCode != http.StatusOK {
		return Manga{}, fmt.Errorf("error: received status code %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Manga{}, err
	}
	log.Printf("API response body: %s", body)

	var mangaData struct {
		Data Manga `json:"data"`
	}

	err = json.Unmarshal(body, &mangaData)
	if err != nil {
		return Manga{}, err
	}

	// Debug: Ensure data is correctly fetched
	log.Printf("Fetched manga: %+v\n", mangaData.Data)

	return mangaData.Data, nil
}
