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
	*AuthService
}
type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
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
		handler.AuthService.Login(body.Email, body.Password)
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
		handler.AuthService.Register(body.Email, body.Password, body.Name)
		res.Json(w, body, 201)
		fmt.Println(body)
	}
}
