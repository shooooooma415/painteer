package v1

import (
	"painteer/handler/v1"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

func InitGroupRoutes(e *echo.Echo, groupService service.GroupService, authService service.AuthService) {

	e.Group("/group")

	e.POST("/", v1.RegisterGroup(groupService))
	e.PUT("/member", v1.JoinGroup(groupService))
	e.GET("/user", v1.GetUserGroup(groupService))
	e.GET("/member", v1.GetGroupMembers(groupService,authService))
	e.GET("/", v1.GetGroup(groupService))
}
