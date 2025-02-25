// package main

// import (
// 	"log"
// 	userPostgresql"painteer/repository/auth/postgresql"
// 	postPostgresql"painteer/repository/posting/postgresql"
// 	groupPostgresql"painteer/repository/group/postgresql"

// 	"painteer/repository/utils"
// 	"painteer/router/v1"
// 	"painteer/service"

// 	"github.com/labstack/echo/v4"
// 	_ "github.com/lib/pq"
// )

// func main() {
// 	e := echo.New()

// 	db, err := utils.ConnectDB()
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}
// 	defer db.Close()

// 	usersRepo := userPostgresql.NewAuthRepository(db)
// 	postsRepo := postPostgresql.NewPostRepository(db)
// 	groupsRepo := groupPostgresql.NewGroupRepository(db)

// 	authService := service.NewAuthService(usersRepo)
// 	postingService := service.NewPostingService(postsRepo)
// 	groupService := service.NewGroupService(groupsRepo)

// 	v1.InitAuthRoutes(e, authService)
// 	v1.InitGroupRoutes(e, groupService,authService)
// 	v1.InitPostingRoutes(e,postingService,groupService,authService)

// 	log.Println("Starting server on :8080...")
// 	e.Logger.Fatal(e.Start(":8080"))
// }
