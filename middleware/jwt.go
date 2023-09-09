package middleware

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(id int, roles string) (string, error) {
	claims := domain.AuthClaims{
		ID:    id,
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Kampung Budaya 2023",
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 3)},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if bearer == "" {
			response.Fail(w, http.StatusUnauthorized, "Missing Authorization Header")
			return
		}

		bearer = strings.ReplaceAll(bearer, "Bearer ", "")

		var claims *domain.AuthClaims
		token, err := jwt.ParseWithClaims(bearer, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET,KEY")), nil
		})
		if err != nil {
			response.Fail(w, http.StatusUnauthorized, "Token Has Been Tampered")
			return
		}

		if !token.Valid {
			response.Fail(w, http.StatusUnauthorized, "Token No Longer Valid")
			return
		}

		r.Header.Set("id", strconv.Itoa(claims.ID))
		r.Header.Set("roles", claims.Roles)

		next.ServeHTTP(w, r)
	})
}
