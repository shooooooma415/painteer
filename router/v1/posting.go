// package v1

// import (
// 	"painteer/handler"
// 	"painteer/service"

// 	"github.com/labstack/echo/v4"
// )

// func InitPostingRoutes(e *echo.Echo, postingService service.PostingService, groupService service.GroupService, authService service.AuthService) {
// 	e.POST("/post", handler.UploadPost(postingService, groupService))
// 	e.GET("/post/map", handler.GetPosts(postingService))
// 	e.GET("/post", handler.GetPost(postingService, authService))
// }
