package auth

import (
	"fmt"
	"net/http"

	"github.com/BibikovAnton/finance-tracker-api/configs"
	"github.com/BibikovAnton/finance-tracker-api/internal/req"
	"github.com/BibikovAnton/finance-tracker-api/internal/res"
)

type AuthHandler struct {
	*configs.Config
}
type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandlerBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		res.Json(w, body, 201)
		fmt.Println(body)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandlerBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		res.Json(w, body, 201)
		fmt.Println(body)
	}
}
