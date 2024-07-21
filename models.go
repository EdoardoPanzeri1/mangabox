package main

import "database/sql"

// Config holds the database and Jikan configurations
type Config struct {
	DB      *sql.DB
	BaseURL string
}

// JikanRequestParams holds the parameters for a Jikan API request
type JikanRequestParams struct {
	ID int `json:id`
}

type Manga struct {
	MalID          int             `json:"mal_id"`
	URL            string          `json:"url"`
	Images         Images          `json:"images"`
	Approved       bool            `json:"approved"`
	Titles         []Title         `json:"titles"`
	Title          string          `json:"title"`
	TitleEnglish   string          `json:"title_english"`
	TitleJapanese  string          `json:"title_japanese"`
	TitleSynonyms  []string        `json:"title_synonyms"`
	Type           string          `json:"type"`
	Chapters       int             `json:"chapters"`
	Volumes        int             `json:"volumes"`
	Status         string          `json:"status"`
	Publishing     bool            `json:"publishing"`
	Published      Published       `json:"published"`
	Score          float32         `json:"score"`
	ScoredBy       int             `json:"scored_by"`
	Rank           int             `json:"rank"`
	Popularity     int             `json:"popularity"`
	Members        int             `json:"members"`
	Favorites      int             `json:"favorites"`
	Synopsis       string          `json:"synopsis"`
	Background     string          `json:"background"`
	Authors        []Author        `json:"authors"`
	Serializations []Serialization `json:"serializations"`
	Genres         []Genre         `json:"genres"`
}

type Images struct {
	JPG  ImageURLs `json:"jpg"`
	WebP ImageURLs `json:"webp"`
}

type ImageURLs struct {
	ImageURL      string `json:"image_url"`
	SmallImageURL string `json:"small_image_url"`
	LargeImageURL string `json:"large_image_url"`
}

type Title struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

type Published struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Prop   Prop   `json:"prop"`
	String string `json:"string"`
}

type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
}

type Prop struct {
	From Date `json:"from"`
	To   Date `json:"to"`
}

type Author struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Serialization struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Genre struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type APIResponse struct {
	Data Manga `json:"data"`
}

type TManga struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	ImageURL string `json:"image_url"`
}
