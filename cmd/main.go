package main

import (
	"fmt"
	"net/http"

	"github.com/BibikovAnton/finance-tracker-api/configs"
	"github.com/BibikovAnton/finance-tracker-api/internal/auth"
	"github.com/BibikovAnton/finance-tracker-api/internal/scors"
	"github.com/BibikovAnton/finance-tracker-api/internal/user"
	"github.com/BibikovAnton/finance-tracker-api/pkg/db"
	"github.com/BibikovAnton/finance-tracker-api/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.ServeMux{}

	//Repository
	userRepository := user.NewUserRepository(db)
	accountsRepositpry := scors.NewAccountsRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)

	//Handlers
	auth.NewAuthHandler(&router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	scors.NewScoreHandler(&router, scors.AccountsHandlerDeps{
		Config:             conf,
		AccountsRepository: accountsRepositpry,
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: middleware.CORS(&router),
	}
	fmt.Println("Server start on port 8000:")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
