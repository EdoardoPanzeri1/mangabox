package main

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt"
)

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
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	ImageURL string `json:"image_url"`
}

type RetrieveManga struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Status      string   `json:"status"`
	CoverArtUrl string   `json:"cover_art_url"`
	IssueNumber int      `json:"issue_number"`
}

type MangaRequest struct {
	ID              string          `json:"id"`
	Status          interface{}     `json:"status"`
	UserID          int32           `json:"user_id"`
	Title           string          `json:"title"`
	IssueNumber     int32           `json:"issue_number"`
	PublicationDate time.Time       `json:"publication_date"`
	Storyline       string          `json:"storyline"`
	CoverArtUrl     string          `json:"cover_art_url"`
	Images          json.RawMessage `json:"images"`
	Authors         json.RawMessage `json:"authors"`
	Serializations  json.RawMessage `json:"serializations"`
	Genres          json.RawMessage `json:"genres"`
	ExplicitGenres  json.RawMessage `json:"explicit_genres"`
	Themes          json.RawMessage `json:"themes"`
	Demographics    json.RawMessage `json:"demographics"`
	Score           float64         `json:"score"`
	ScoredBy        int32           `json:"scored_by"`
	Rank            int32           `json:"rank"`
	Popularity      int32           `json:"popularity"`
	Members         int32           `json:"members"`
	Favorites       int32           `json:"favorites"`
	Synopsis        string          `json:"synopsis"`
	Background      string          `json:"background"`
	Relations       json.RawMessage `json:"relations"`
	ExternalLinks   json.RawMessage `json:"external_links"`
}

type UpdateStatusRequest struct {
	UserID int32  `json:"user_id"`
	Status string `json:"status"`
}

// User models
type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UpdateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
