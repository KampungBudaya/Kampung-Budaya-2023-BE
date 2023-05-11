package http

import (
	"context"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/contest/usecase"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/middleware"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
	"github.com/gorilla/mux"
)

type ContestHandler struct {
	contest usecase.ContestUsecaseImpl
	router  *mux.Router
}

func NewContestHandler(router *mux.Router, Contest usecase.ContestUsecaseImpl) {
	compHandler := ContestHandler{
		contest: Contest,
		router:  router,
	}
	compHandler.router.HandleFunc("/contest", compHandler.RegisterContest).Methods(http.MethodPost)

	participantsHandler := compHandler.router.PathPrefix("/participants").Subrouter()
	participantsHandler.Use(middleware.ValidateJWT)
	participantsHandler.HandleFunc("/participants", compHandler.GetAllParticipants).Methods(http.MethodGet)
	participantsHandler.HandleFunc("/participants/{id}", compHandler.GetParticipantByID).Methods(http.MethodGet)
	participantsHandler.HandleFunc("/participants/{id}", compHandler.AcceptParticipant).Methods(http.MethodPatch)
}

func (h *ContestHandler) RegisterContest(w http.ResponseWriter, r *http.Request) {
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
		user := r.Context().Value("user").(domain.UserContext)
		if !strings.Contains(user.Roles, "Super Admin") {
			code = http.StatusUnauthorized
			errChan <- err
			return
		}

		contestID, err := strconv.Atoi(r.FormValue("contestID"))
		if err != nil {
			code = http.StatusBadRequest
			errChan <- err
			return
		}
		name := r.FormValue("name")
		origin := r.FormValue("origin")
		phoneNumber := r.FormValue("phoneNumber")
		videoURL := r.FormValue("videoURL")

		req := domain.StoreParticipant{
			ContestID:   contestID,
			Name:        name,
			Origin:      origin,
			PhoneNumber: phoneNumber,
			VideoURL:    videoURL,
		}

		form, _, err := r.FormFile("form")
		if err != nil {
			code = http.StatusBadRequest
			errChan <- err
			return
		}
		paymentProof, _, err := r.FormFile("paymentProof")
		if err != nil {
			code = http.StatusBadRequest
			errChan <- err
			return
		}

		res, err := h.contest.RegisterContest(ctx, req, []multipart.File{form, paymentProof})
		if err != nil {
			code = http.StatusInternalServerError
			errChan <- err
			return
		}
		resChan <- res
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

func (h *ContestHandler) GetAllParticipants(w http.ResponseWriter, r *http.Request) {
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
		res, err := h.contest.GetAllParticipants(ctx)
		if err != nil {
			code = http.StatusInternalServerError
			errChan <- err
			return
		}
		resChan <- res
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

func (h *ContestHandler) GetParticipantByID(w http.ResponseWriter, r *http.Request) {
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
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			code = http.StatusBadRequest
			errChan <- err
			return
		}

		res, err := h.contest.GetParticipantByID(ctx, id)
		if err != nil {
			code = http.StatusInternalServerError
			errChan <- err
			return
		}
		resChan <- res
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

func (h *ContestHandler) AcceptParticipant(w http.ResponseWriter, r *http.Request) {
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

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		code = http.StatusBadRequest
		return
	}

	err = h.contest.AcceptParticipant(ctx, id)
	if err != nil {
		code = http.StatusInternalServerError
	}

	select {
	case <-ctx.Done():
		err = ctx.Err()
		code = http.StatusRequestTimeout
		return
	default:
		data = "SUCCES ACCEPT PARTICIPANT"
	}
}
