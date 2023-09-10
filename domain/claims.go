package domain

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	ID    int    `json:"id"`
	Roles string `json:"roles"`
	jwt.RegisteredClaims
}

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	jwt.StandardClaims
}
