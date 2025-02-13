// package v1

// import (
// 	"painteer/handler"
// 	"painteer/service"

// 	"github.com/labstack/echo/v4"
// )

// func InitGroupRoutes(e *echo.Echo, groupService service.GroupService, authService service.AuthService) {
// 	e.POST("/group", handler.RegisterGroup(groupService))
// 	e.PUT("/group/member", handler.JoinGroup(groupService))
// 	e.GET("/group/user", handler.GetUserGroup(groupService))
// 	e.GET("/group/member", handler.GetGroupMembers(groupService,authService))
// 	e.GET("/group", handler.GetGroup(groupService))
// }
