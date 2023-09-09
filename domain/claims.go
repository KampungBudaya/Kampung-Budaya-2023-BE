package domain

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	ID    int    `json:"id"`
	Roles string `json:"roles"`
	jwt.RegisteredClaims
}

type GoogleClaims struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	jwt.StandardClaims
}
