package scors

import (
	"net/http"
	"strconv"

	"github.com/BibikovAnton/finance-tracker-api/configs"
	"github.com/BibikovAnton/finance-tracker-api/internal/req"
	"github.com/BibikovAnton/finance-tracker-api/internal/res"
	"github.com/BibikovAnton/finance-tracker-api/pkg/middleware"
	"gorm.io/gorm"
)

type AccountsHandlerDeps struct {
	AccountsRepository *AccountsRepository
	Config             *configs.Config
}

type AccountsHandler struct {
	AccountsRepository *AccountsRepository
}

func NewScoreHandler(router *http.ServeMux, deps AccountsHandlerDeps) {
	handler := &AccountsHandler{
		AccountsRepository: deps.AccountsRepository,
	}

	router.Handle("PATCH /accounts/{id}", middleware.IsAuth(handler.Update()))
	router.HandleFunc("POST /accounts", handler.Create())
	router.HandleFunc("DELETE /accounts/{id}", handler.Delete())
	router.HandleFunc("GET /accounts/{id}", handler.GetById())
	router.HandleFunc("GET /accounts", handler.Get())
}

func (handler *AccountsHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandlerBody[AccountsCreateRequest](&w, r)
		if err != nil {
			return
		}
		account := NewAccount(
			body.Name,
			body.Type,
			body.Balance,
			body.Currency,
		)

		createdAccount, err := handler.AccountsRepository.Create(account)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdAccount, 201)
	}

}

func (handler *AccountsHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandlerBody[AccountsUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseInt(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		account, err := handler.AccountsRepository.Update(&Account{
			Model:   gorm.Model{ID: uint(id)},
			Name:    body.Name,
			Type:    body.Type,
			Balance: body.Balance,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, account, 201)
	}
}

func (handler *AccountsHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseInt(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handler.AccountsRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, 200)

	}
}

func (handler *AccountsHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var accounts []Account
		data, err := handler.AccountsRepository.GetAll(&accounts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, data, 200)

	}
}

func (handler *AccountsHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		data, err := handler.AccountsRepository.Get(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, data, 200)

	}
}
