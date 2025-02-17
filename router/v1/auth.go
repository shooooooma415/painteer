package v1

import (
	"painteer/handler"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(e *echo.Echo, authService service.AuthService) {
	authGroup := e.Group("/auth")

	authGroup.POST("/signup", handler.SignUp(authService))
	authGroup.GET("/signin", handler.SignIn(authService))
	authGroup.GET("/profile", handler.GetUserByID(authService))
}
