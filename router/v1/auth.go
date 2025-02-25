package v1

import (
	"painteer/handler/v1"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(e *echo.Echo, authService service.AuthService) {
	authGroup := e.Group("/auth")

	authGroup.POST("/signup", v1.SignUp(authService))
	authGroup.GET("/signin", v1.SignIn(authService))
	authGroup.GET("/profile", v1.GetUserByID(authService))
}
