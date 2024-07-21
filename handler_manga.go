package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
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
	w.Header().Set("COntent-Type", "applications/json")
	json.NewEncoder(w).Encode(manga)
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

func (cfg *apiConfig) handlerGetMangaDB(w http.ResponseWriter, r *http.Request) {
	// Fetching manga ID from URL parameters
	mangaIDSrt := r.URL.Query().Get("id")
	if mangaIDSrt == "" {
		respondWithError(w, http.StatusBadRequest, "id parameter is missing")
		return
	}

	mangaID, err := strconv.Atoi(mangaIDSrt)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	//Fetch manga details from the database using sqlc generated code
	manga, err := s.db.GetMangaById(context.Background(), mangaID)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "manga not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
	}
	return

	w.Header().Set("Content-Type", "applications/json")
	json.NewEncoder(w).Encode(manga)

}
