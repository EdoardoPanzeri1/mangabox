package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/EdoardoPanzeri1/mangabox/internal/database"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) handlerRegistration(w http.ResponseWriter, r *http.Request) {
	if !checkRequest(w, r) {
		return
	}

	// Decode the incoming JSON request
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the request data
	if req.Username == "" || req.Email == "" || req.Password == "" {
		respondWithError(w, http.StatusBadRequest, "All fields are required")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Prepare the database parameters
	params := database.CreateUserParams{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	// Insert the new user into the database
	if err := cfg.DB.CreateUser(r.Context(), params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "User registered successfully"})
}

var jwtKey []byte

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	if !checkRequest(w, r) {
		return
	}

	// Decode the incoming JSON request
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the request data
	if req.Username == "" || req.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Both username and password are required")
		return
	}

	// Fetch the user from the database using the username
	user, err := cfg.DB.FetchUserByUsername(r.Context(), req.Username)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// Compare the provided password with the hashed password stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// Generate a JWT token upon successful authentication
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	respondWithJSON(w, http.StatusOK, LoginResponse{Token: tokenString})
}

func checkRequest(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return false
	}

	if r.Header.Get("Content-Type") != "application/json" {
		respondWithError(w, http.StatusUnsupportedMediaType, "Content Type is not application/json")
		return false
	}

	return true
}
