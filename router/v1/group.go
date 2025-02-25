package v1

import (
	"painteer/handler"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

func InitGroupRoutes(e *echo.Echo, groupService service.GroupService, authService service.AuthService) {

	e.Group("/group")

	e.POST("/", handler.RegisterGroup(groupService))
	e.PUT("/member", handler.JoinGroup(groupService))
	e.GET("/user", handler.GetUserGroup(groupService))
	e.GET("/member", handler.GetGroupMembers(groupService,authService))
	e.GET("/", handler.GetGroup(groupService))
}
