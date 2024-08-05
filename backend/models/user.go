package models

type User struct {
	ID       int
	Email    string
	Password string
}

// DB interaction methods to be added, like CreateUser, FindUserByEmail...
