package main

import (
	"os"

	"github.com/BibikovAnton/finance-tracker-api/internal/scors"
	"github.com/BibikovAnton/finance-tracker-api/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//подключили окружение
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	//подключили базу данных
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{}) //подключение к базе данных, принемает driver
	if err != nil {
		panic(err)
	}
	//авто-мигрируем
	db.AutoMigrate(&user.User{}, scors.Account{})
}
