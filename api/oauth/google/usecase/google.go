package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/repository"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/golang-jwt/jwt/v4"
)

type GoogleUsecaseImpl interface {
	Authenticate(email string, providerID string, ctx context.Context) error
	ValidateGoogleJWT(tokenString string) (domain.GoogleClaims, error)
}

type GoogleUsecase struct {
	oauth    repository.OAuthRepositoryImpl
	clientID string
}

func NewGoogleUsecase(oauth repository.OAuthRepositoryImpl, clientID string) GoogleUsecaseImpl {
	return &GoogleUsecase{
		oauth:    oauth,
		clientID: clientID,
	}
}

func (gu *GoogleUsecase) Authenticate(email string, providerID string, ctx context.Context) error {
	tx, err := gu.oauth.BeginTx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	user, err := gu.oauth.GetByEmail(email, ctx)
	if err != nil {
		return err
	}

	if user.ProviderID == "" {
		if err = gu.oauth.UpdateProviderID(user.ID, providerID, ctx); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (gu *GoogleUsecase) ValidateGoogleJWT(tokenArg string) (domain.GoogleClaims, error) {
	var claimsStruct domain.GoogleClaims

	token, err := jwt.ParseWithClaims(tokenArg, &claimsStruct, func(t *jwt.Token) (interface{}, error) {
		pem, err := getGooglePublicKey(fmt.Sprintf("%s", t.Header["kid"]))
		if err != nil {
			return nil, err
		}
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, err
		}

		return key, nil
	})
	if err != nil {
		return domain.GoogleClaims{}, err
	}

	claims, exists := token.Claims.(*domain.GoogleClaims)
	if !exists {
		return domain.GoogleClaims{}, errors.New("Invalid Google JWT")
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return domain.GoogleClaims{}, errors.New("iss is invalid")
	}

	if claims.Audience != gu.clientID {
		return domain.GoogleClaims{}, errors.New("aud is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return domain.GoogleClaims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}

func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	data := map[string]string{}
	if err = json.Unmarshal(raw, &data); err != nil {
		return "", err
	}
	key, exists := data[keyID]
	if !exists {
		return "", errors.New("Key does not exists")
	}

	return key, nil
}
