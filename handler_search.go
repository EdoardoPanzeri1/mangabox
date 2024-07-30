package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func (cfg *apiConfig) handlerSearchManga(w http.ResponseWriter, r *http.Request) {
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

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Error decoding response from Jikan API: %v\n", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to decode responde from Jikan API")
		return
	}

	dataRaw, ok := result["data"].([]interface{})
	if !ok {
		respondWithError(w, http.StatusInternalServerError, "Failed to decode response for Jikan")
		return
	}

	var data []map[string]interface{}
	for _, item := range dataRaw {
		itemMap, ok := item.(map[string]interface{})
		if ok {
			data = append(data, itemMap)
		}
	}

	transformedResult := transformResult(data)
	respondWithJSON(w, http.StatusOK, transformedResult)
}

func transformResult(rawData []map[string]interface{}) []TManga {
	var transformedResults []TManga

	for _, item := range rawData {
		title, exists := item["title"].(string)
		if !exists {
			continue
		}

		// Check the nested structure for authors
		authors := "Unknown Author"
		if authorData, exists := item["authors"].([]interface{}); exists && len(authorData) > 0 {
			if auth, ok := authorData[0].(map[string]interface{}); ok {
				authors, _ = auth["name"].(string)
			}
		}

		// Might need to check a different field for image URL
		imageURL, exists := item["images"].(map[string]interface{})
		var imageLink string
		if exists {
			if jpg, ok := imageURL["jpg"].(map[string]interface{}); ok {
				imageLink, _ = jpg["image_url"].(string)
			}
		}
		if imageLink == "" {
			imageLink = "default_image_url.jpg"
		}

		transformedResults = append(transformedResults, TManga{
			Title:    title,
			Author:   authors,
			ImageURL: imageLink,
		})
	}

	return transformedResults
}
