package auth

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
	GetUserByID(c echo.Context) error
}
