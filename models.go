package main

import "database/sql"

// Database and ComicWiz configuration
type Config struct {
	DB      *sql.DB
	BaseURL string
	APIKey  string
}

// Parameters for a ComicWiz API request
type ComicVineRequestParams struct {
	Format    string
	FieldList string
	Limit     int
	Offset    int
	Sort      string
	Filter    string
}

type ComicVineVolume struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FirstIssue  string `json:"first_issue"`
	LastIssue   string `json:"last_issue"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
}

type ComicVineResponse struct {
	Results []ComicVineVolume `json:"result"`
}
