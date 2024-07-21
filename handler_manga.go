package main

import (
	"encoding/json"
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
	response, err := http.Get("http://api.jikan.moe/v3/manga/" + strconv.Itoa(mangaID))
	if err != nil {
		return Manga{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Manga{}, err
	}

	var manga Manga
	err = json.NewDecoder(response.Body).Decode(&manga)
	if err != nil {
		return Manga{}, err
	}

	return manga, nil
}
