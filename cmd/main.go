package main

import (
	"fmt"
	"net/http"

	"github.com/BibikovAnton/finance-tracker-api/configs"
	"github.com/BibikovAnton/finance-tracker-api/internal/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.ServeMux{}

	//Handlers
	auth.NewAuthHandler(&router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: &router,
	}
	fmt.Println("Server start on port 8000:")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
