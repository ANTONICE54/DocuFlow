package models

import "time"

type User struct {
	ID             uint
	Name           string
	Surname        string
	Email          string
	Country        string
	HashedPassword string
	CreatedAt      time.Time
}

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Country  string `json:"country"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	Token string `json:"token"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

type VerifyRequest struct {
	Token string `json:"token"`
}

type VerifyResponse struct {
	UserID uint `json:"user_id"`
}

type UpdateUserRequest struct {
	ID      uint
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Country string `json:"country"`
}

type UpdateUserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Country string `json:"country"`
}

type GetUserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Country string `json:"country"`
}
