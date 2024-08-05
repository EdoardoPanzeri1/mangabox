package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/EdoardoPanzeri1/mangabox/internal/database"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sqlc-dev/pqtype"
)

func (cfg *apiConfig) handlerRetrieveCatalog(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	nullUserID := stringToNullUUID(userID)

	// Debugging
	log.Printf("Retrieve catalog from user_id: %s", userID)

	ctx := r.Context()
	log.Println("handlerAddToCatalog: Request context obtained") // Debugging

	// Use the generated RetrieveCatalog method
	rows, err := cfg.DB.RetrieveCatalog(ctx, nullUserID)
	if err != nil {
		// Debugging
		log.Printf("Error retrieving catalog from database %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve the catalog")
		return
	}

	var mangas []RetrieveManga
	for _, row := range rows {
		var manga RetrieveManga

		// Handle pqtype.NullRawMessage
		if row.Authors.Valid {
			// Unmarshal the authors JSON string into the Authors slice
			var authorsString string
			err := json.Unmarshal(row.Authors.RawMessage, &authorsString)
			if err != nil {
				// Debugging
				log.Printf("Raw Authors JSON: %s", string(row.Authors.RawMessage))
				log.Printf("Error parsing authors from JSON: %v", err)
				respondWithError(w, http.StatusInternalServerError, "Failed to parse authors")
				return
			}

			if err := json.Unmarshal([]byte(authorsString), &manga.Authors); err != nil {
				log.Printf("Error parsing author from JSON: %v", err)
				respondWithError(w, http.StatusInternalServerError, "Failed to parse authors")
				return
			}
		} else {
			manga.Authors = []string{}
		}

		if row.Status != nil {
			// Convert []uint8 to string
			statusByte, ok := row.Status.([]uint8)
			if ok {
				manga.Status = string(statusByte)
			} else {
				log.Printf("Error: Status is not a string, it's %T", row.Status)
				respondWithError(w, http.StatusInternalServerError, "Invalid Status type")
				return
			}
		} else {
			manga.Status = ""
		}

		if row.CoverArtUrl.Valid {
			manga.CoverArtUrl = row.CoverArtUrl.String
		} else {
			manga.CoverArtUrl = ""
		}

		manga.IssueNumber = int(row.IssueNumber)

		manga.ID = row.ID
		manga.Title = row.Title
		mangas = append(mangas, manga)
	}

	respondWithJSON(w, http.StatusOK, mangas)
}

func (cfg *apiConfig) handlerAddToCatalog(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON request into the MangaRequest struct
	var req MangaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("handlerAddToCatalog: Invalid request payload: %v\n", err) // Debugging
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	log.Printf("handlerAddToCatalog: Request payload: %+v\n", req)

	// Get the request context
	ctx := r.Context()
	log.Println("handlerAddToCatalog: Request context obtained") // Debugging

	nullUserID := stringToNullUUID(req.UserID)

	// Create an instance of InsertMangaIntoCatalogParams with the decode data
	params := database.InsertMangaIntoCatalogParams{
		ID:              req.ID,
		Status:          req.Status,
		UserID:          nullUserID,
		Title:           req.Title,
		IssueNumber:     req.IssueNumber,
		PublicationDate: req.PublicationDate,
		Storyline:       sql.NullString{String: req.Storyline, Valid: true},
		CoverArtUrl:     sql.NullString{String: req.CoverArtUrl, Valid: true},
		UpdatedAt:       sql.NullTime{Time: time.Now(), Valid: true},
		Images:          pqtype.NullRawMessage{RawMessage: req.Images, Valid: true},
		Authors:         pqtype.NullRawMessage{RawMessage: req.Authors, Valid: true},
		Serializations:  pqtype.NullRawMessage{RawMessage: req.Serializations, Valid: true},
		Genres:          pqtype.NullRawMessage{RawMessage: req.Genres, Valid: true},
		ExplicitGenres:  pqtype.NullRawMessage{RawMessage: req.ExplicitGenres, Valid: true},
		Themes:          pqtype.NullRawMessage{RawMessage: req.Themes, Valid: true},
		Demographics:    pqtype.NullRawMessage{RawMessage: req.Demographics, Valid: true},
		Score:           sql.NullFloat64{Float64: req.Score, Valid: true},
		ScoredBy:        sql.NullInt32{Int32: req.ScoredBy, Valid: true},
		Rank:            sql.NullInt32{Int32: req.Rank, Valid: true},
		Popularity:      sql.NullInt32{Int32: req.Popularity, Valid: true},
		Members:         sql.NullInt32{Int32: req.Members, Valid: true},
		Favorites:       sql.NullInt32{Int32: req.Favorites, Valid: true},
		Synopsis:        sql.NullString{String: req.Synopsis, Valid: true},
		Background:      sql.NullString{String: req.Background, Valid: true},
		Relations:       pqtype.NullRawMessage{RawMessage: req.Relations, Valid: true},
		ExternalLinks:   pqtype.NullRawMessage{RawMessage: req.ExternalLinks, Valid: true},
	}

	log.Printf("handlerAddToCatalog: Insert parameters: %+v\n", params) // Debugging

	// Call the InsertMangaIntoCatalog method with the constructed parameters
	if err := cfg.DB.InsertMangaIntoCatalog(ctx, params); err != nil {
		log.Printf("handlerAddToCatalog: Error inserting manga into catalog: %v\n", err) // Debugging
		respondWithError(w, http.StatusInternalServerError, "Failed to add manga to catalog")
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Manga added to catalog"})
}

func StringToStatus(s string) Status {
	return Status(s)
}

func (cfg *apiConfig) handlerStatusManga(w http.ResponseWriter, r *http.Request) {
	// Extract the manga ID from the URL
	vars := mux.Vars(r)
	mangaID := vars["id"]
	if mangaID == "" {
		respondWithError(w, http.StatusBadRequest, "Manga ID is required")
		return
	}

	// Decode the incoming JSON request into the UpdateStatusRequest struct
	var req UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Ensure that the status provided is valid
	if req.Status != "read" {
		respondWithError(w, http.StatusBadRequest, "Invalid status value; only 'read' is allowed")
		return
	}

	// Debugg
	log.Printf("Payload received: %+v", req)

	ctx := r.Context()

	nullUserID := stringToNullUUID(req.UserID)

	// Create parameters for the update query
	params := database.UpdateStatusReadParams{
		ID:     mangaID,
		UserID: nullUserID,
	}

	// Call the UpdateStatusRead method with the constructed parameters
	if err := cfg.DB.UpdateStatusRead(ctx, params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update manga status")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Manga status updated to 'read'"})
}

func (cfg *apiConfig) handlerDeleteManga(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract the manga ID from the URL path parameters
	vars := mux.Vars(r)
	mangaID := vars["id"]
	if mangaID == "" {
		respondWithError(w, http.StatusBadRequest, "manga ID is required")
		return
	}

	// Extract user ID from query parameters
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "user id is required")
		return
	}

	ctx := r.Context()

	nullUserID := stringToNullUUID(userID)

	// Create parameters for the delete query
	params := database.DeleteMangaParams{
		ID:     mangaID,
		UserID: nullUserID,
	}

	// Call the deleteManga method with the constructed params
	if err := cfg.DB.DeleteManga(ctx, params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete manga from the catalog")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Manga deleted from the catalog"})
}

// stringToNullUUID converts a string to a uuid.NullUUID
func stringToNullUUID(str string) uuid.NullUUID {
	var nullUUID uuid.NullUUID
	if uid, err := uuid.Parse(str); err == nil {
		nullUUID = uuid.NullUUID{UUID: uid, Valid: true}
	} else {
		nullUUID = uuid.NullUUID{Valid: false}
	}
	return nullUUID
}
