package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/google/usecase"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
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
	googleOAuth.Use(middleware.Guest)
	googleOAuth.HandleFunc("", handler.SignIn).Methods(http.MethodPost)
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

	var request domain.AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		code = http.StatusBadRequest
		return
	}

	claims, err := h.google.ValidateGoogleJWT(request.Token)
	if err != nil {
		code = http.StatusBadRequest
		return
	}

	user, err := h.google.SearchUser(claims.Email, ctx)
	if err != nil {
		code = http.StatusUnauthorized
		return
	}

	token, err := middleware.GenerateJWT(user.ID, user.Roles)
	if err != nil {
		code = http.StatusInternalServerError
		return
	}

	select {
	case <-ctx.Done():
		err = ctx.Err()
		code = http.StatusRequestTimeout
	default:
		response := domain.AuthResponse{SignedToken: token}
		response.Populate(user.Roles)
		data = response
	}
}
