package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/faq/usecase"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
	"github.com/gorilla/mux"
)

type FaqHandler struct {
	router *mux.Router
	faq    *usecase.FaqUsecaseImpl
}

func NewFaqHandler(router *mux.Router, faq *usecase.FaqUsecaseImpl) {
	_ = &FaqHandler{
		router: router,
		faq:    faq,
	}

	// faqHandler.router.HandleFunc("/faq", faqHandler.GetFaq).Methods("GET")
}

func (h *FaqHandler) AddFaq(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
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
		body, err := io.ReadAll(r.Body)
		if err != nil {
			errChan <- err
			code = http.StatusBadRequest
			return
		}

		req := domain.FaqReq{}
		if err := json.Unmarshal(body, &req); err != nil {
			errChan <- err
			code = http.StatusBadRequest
			return
		}

	}()

	select {
	case <-ctx.Done():
		err = ctx.Err()
		code = http.StatusRequestTimeout
	case err = <-errChan:
	case data = <-resChan:

	}
}
