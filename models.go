package main

import "database/sql"

// Config holds the database and ComicVine configurations
type Config struct {
	DB      *sql.DB
	BaseURL string
	APIKey  string
}

// ComicVineRequestParams holds the parameters for a ComicVine API request
type ComicVineRequestParams struct {
	Format    string
	FieldList string
	Limit     int
	Offset    int
	Sort      string
	Filter    string
}

// ComicVineVolume represents a single comic volume from ComicVine
type ComicVineVolume struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FirstIssue  string `json:"first_issue"`
	LastIssue   string `json:"last_issue"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
}

// ComicVineResponse represents the response from ComicVine API
type ComicVineResponse struct {
	Results []ComicVineVolume `json:"result"`
}
