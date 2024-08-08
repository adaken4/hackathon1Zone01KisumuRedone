package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/adaken4/peerfund-p2p-lending/backend/store"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Define a secret key for signing the JWT
var jwtKey = []byte("xkUbdz6r$")

// Create a struct to represent the claims
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// func LoginUser(w http.ResponseWriter, r *http.Request) {
// 	var existingUser models.User
// 	err := json.NewDecoder(r.Body).Decode(&existingUser)
// 	if err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}
// 	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(existingUser.Password)); err != nil {
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("Login successful")
// }

// controllers/auth.go
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("psw")

	user, exists := store.GetUser(email)
	if !exists {
		// Redirect to the registration page if the user does not exist
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create the token
	expirationTime := time.Now().Add(5 * time.Minute) // Token expiration time
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	// Set the token in a cookie (optional)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	// Redirect or respond with success
	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}

// Middleware to validate JWT
func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Set the email in the context for use in the next handler
		ctx := context.WithValue(r.Context(), "email", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
