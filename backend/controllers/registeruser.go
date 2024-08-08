package controllers

import (
	"fmt"
	"net/http"

	"github.com/adaken4/peerfund-p2p-lending/backend/models"
	"github.com/adaken4/peerfund-p2p-lending/backend/store"
	"github.com/adaken4/peerfund-p2p-lending/backend/utils"
)

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
