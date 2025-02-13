// package v1

// import (
// 	"painteer/handler"
// 	"painteer/service"

// 	"github.com/labstack/echo/v4"
// )

// func InitAuthRoutes(e *echo.Echo, authService service.AuthService) {
// 	e.POST("/auth/signup", handler.SignUp(authService))
// 	e.GET("/auth/signin", handler.SignIn(authService))
// 	e.GET("/auth/profile", handler.GetUserByID(authService))
// }
