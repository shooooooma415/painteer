package v1

import (
	"painteer/handler/v1"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

func InitPostingRoutes(e *echo.Echo, postingService service.PostingService, groupService service.GroupService, authService service.AuthService) {
	e.POST("/post", v1.UploadPost(postingService))
	e.GET("/post/map", v1.GetPosts(postingService))
	e.GET("/post", v1.GetPost(postingService, authService))
}
