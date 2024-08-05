package store

import (
	"fmt"
	"sync"

	"github.com/adaken4/peerfund-p2p-lending/backend/models"
)

var (
	userStore  = make(map[string]models.User) // Map to store users by email
	mu         sync.Mutex                     // Mutex to protect access to the user store
	nextUserID = 1                            // Counter for the next user ID
)

// AddUser adds a new user to the in-memory store
func AddUser(user models.User) {
	user.ID = nextUserID
	mu.Lock()
	defer mu.Unlock()

	userStore[user.Email] = user
	fmt.Println(userStore)
	nextUserID++
}

// GetUser retrieves a user from the in-memory store by email
func GetUser(email string) (models.User, bool) {
	mu.Lock()
	defer mu.Unlock()

	user, exists := userStore[email]
	return user, exists
}
