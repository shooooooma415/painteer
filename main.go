package main

import (
	"log"
	"os"
	userPostgresql "painteer/repository/auth/postgresql"

	// postPostgresql"painteer/repository/post/postgresql"
	groupPostgresql "painteer/repository/group/postgresql"

	"painteer/repository/utils"
	v1 "painteer/router/v1"
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

	usersRepo := userPostgresql.NewAuthRepository(db)
	// postsRepo := postPostgresql.NewPostRepository(db)
	groupsRepo := groupPostgresql.NewGroupRepository(db)

	authService := service.NewAuthService(usersRepo)
	// postingService := service.NewPostingService(postsRepo)
	groupService := service.NewGroupService(groupsRepo)

	v1.InitAuthRoutes(e, authService)
	v1.InitGroupRoutes(e, groupService, authService)
	// v1.InitPostingRoutes(e,postingService,groupService,authService)

	port := os.Getenv("RUN_PORT")

	log.Printf("Starting server on :%s...", port)
	e.Logger.Fatal(e.Start(":" + port))
}
