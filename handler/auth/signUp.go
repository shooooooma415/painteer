package handler

import (
	"net/http"
	"painteer/model"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type signUpResponse struct {
	UserId model.UserId `json:"user_id"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{Error: message}
}

func SignUp(authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.CreateUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		}

		user, err := authService.RegisterUser(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		}

		response := signUpResponse{
			UserId: user.UserId,
		}

		return c.JSON(http.StatusOK, response)
	}
}
