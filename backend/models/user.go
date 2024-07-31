package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// DB interaction methods to be added, like CreateUser, FindUserByEmail...
