package main

import "database/sql"

// Config holds the database and Jikan configurations
type Config struct {
	DB      *sql.DB
	BaseURL string
}

// JikanRequestParams holds the parameters for a Jikan API request
type JikanEequestParams struct {
	ID int `json:id`
}

type Author struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Image struct {
	ImageURL      string `json:"image_url"`
	SmallImageURL string `json:"small_image_url"`
	LargeImageURL string `json:"large_image_url"`
}

type Title struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

type Published struct {
	From string `json:"from"`
	To   string `json:"to"`
	Prop struct {
		From struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"from"`
		To struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"to"`
		String string `json:"string"`
	} `json:"prop"`
}

type Genre struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type RelationEntry struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Relation struct {
	Relation string          `json:"relation"`
	Entry    []RelationEntry `json:"entry"`
}

type ExternalLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Manga struct {
	MalID          int              `json:"mal_id"`
	URL            string           `json:"url"`
	Images         map[string]Image `json:"images"` // Images in both jpg and webp formats
	Approved       bool             `json:"approved"`
	Titles         []Title          `json:"titles"`
	Title          string           `json:"title"`
	TitleEnglish   string           `json:"title_english"`
	TitleJapanese  string           `json:"title_japanese"`
	TitleSynonyms  []string         `json:"title_synonyms"`
	Type           string           `json:"type"`
	Chapters       int              `json:"chapters"`
	Volumes        int              `json:"volumes"`
	Status         string           `json:"status"`
	Publishing     bool             `json:"publishing"`
	Published      Published        `json:"published"`
	Score          float64          `json:"score"`
	ScoredBy       int              `json:"scored_by"`
	Rank           int              `json:"rank"`
	Popularity     int              `json:"popularity"`
	Members        int              `json:"members"`
	Favorites      int              `json:"favorites"`
	Synopsis       string           `json:"synopsis"`
	Background     string           `json:"background"`
	Authors        []Author         `json:"authors"`
	Serializations []Genre          `json:"serializations"`
	Genres         []Genre          `json:"genres"`
	ExplicitGenres []Genre          `json:"explicit_genres"`
	Themes         []Genre          `json:"themes"`
	Demographics   []Genre          `json:"demographics"`
	Relations      []Relation       `json:"relations"`
	External       []ExternalLink   `json:"external"`
}

type APIResponse struct {
	Data Manga `json:"data"`
}

type TManga struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	ImageURL string `json:"image_url"`
}
