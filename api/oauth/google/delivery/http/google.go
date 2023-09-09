package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/google/usecase"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/middleware"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
	"github.com/gorilla/mux"
)

type GoogleHandler struct {
	google usecase.GoogleUsecaseImpl
	router *mux.Router
}

func NewGoogleHandler(router *mux.Router, google usecase.GoogleUsecaseImpl) {
	handler := GoogleHandler{
		google: google,
		router: router,
	}

	googleOAuth := handler.router.PathPrefix("/oauth/google").Subrouter()
	googleOAuth.HandleFunc("/", handler.SignIn)
	googleOAuth.Use(middleware.Guest)
}

func (h *GoogleHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	var (
		err  error
		code = http.StatusOK
		data interface{}
	)

	defer func() {
		if err != nil {
			response.Fail(w, code, err.Error())
			return
		}
		response.Success(w, code, data)
	}()

	errChan := make(chan error, 1)
	resChan := make(chan interface{}, 1)

	go func() {
		var request struct{ token *string }

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			code = http.StatusBadRequest
			errChan <- err
			return
		}

		claims, err := h.google.ValidateGoogleJWT(*request.token)
		if err != nil {
			code = http.StatusBadRequest
			errChan <- err
			return
		}

		user, err := h.google.SearchUser(claims.Email, ctx)
		if err != nil {
			code = http.StatusUnauthorized
			errChan <- err
			return
		}

		if user.ProviderID == "" {
			if h.google.UpdateProviderID(user.ID, claims.ID, ctx); err != nil {
				code = http.StatusInternalServerError
				errChan <- err
				return
			}
		}

		token, err := middleware.GenerateJWT(user.ID, user.Roles)
		if err != nil {
			code = http.StatusInternalServerError
			errChan <- err
			return
		}

		resChan <- map[string]string{
			"token": token,
		}
	}()

	select {
	case <-ctx.Done():
		err = ctx.Err()
		code = http.StatusRequestTimeout
		return
	case err = <-errChan:
	case data = <-resChan:
	}
}
