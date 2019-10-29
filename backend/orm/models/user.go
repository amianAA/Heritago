package models

import "github.com/dgrijalva/jwt-go"

// User Model
type User struct {
	BaseModelSoftDelete
	Email               string  `gorm:"not null;unique_index:idx_email"`
	UserID              *string
	FirstName           *string
	LastName            *string
	Country             *string
	Password            *string
}

// JWT User Utils
type UserAuth struct {
   UserID    string
   Roles     []string
   IPAddress string
   Token     string
}

type UserClaims struct {
   jwt.StandardClaims
   UserId string `json:"user_id"`
}