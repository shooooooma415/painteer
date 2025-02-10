package main

import (
	"log"
	"painteer/repository/auth/postgresql"
	"painteer/repository/utils"
	"painteer/router/v1"
	"painteer/service"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	usersRepo := postgresql.NewAuthRepository(db)

	authService := service.NewAuthService(usersRepo)

	v1.InitAuthRoutes(e, authService)

	log.Println("Starting server on :8080...")
	e.Logger.Fatal(e.Start(":8080"))
}
