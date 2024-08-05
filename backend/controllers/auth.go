package controllers

import (
	"fmt"
	"net/http"

	"github.com/adaken4/peerfund-p2p-lending/backend/models"
	"github.com/adaken4/peerfund-p2p-lending/backend/store"
	"github.com/adaken4/peerfund-p2p-lending/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

// func RegisterUser(w http.ResponseWriter, r *http.Request) {
// 	var user models.User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	hashedPassword, err := utils.HashPassword(user.Password)
// 	if err != nil {
// 		http.Error(w, "Error hashing password", http.StatusInternalServerError)
// 		return
// 	}
// 	user.Password = hashedPassword

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(user)
// }

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Extract form values
	email := r.FormValue("email")
	password := r.FormValue("psw")

	fmt.Println(r.Form)
	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}

	// Create user object
	user := models.User{
		Email:    email,
		Password: hashedPassword,
	}
	fmt.Println(user)

	// Store user in the in-memory map
	store.AddUser(user)

	// Here you would typically save the user to the database
	// For now, we'll just simulate a successful registration

	// Redirect to a success page or login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
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
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Successful login logic (e.g., set session, redirect, etc.)
	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}
