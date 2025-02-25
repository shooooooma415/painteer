package handler

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error string `json:"error"`
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

		return c.JSON(http.StatusOK, user)
	}
}

func SignIn(authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		authIdStr := c.QueryParam("auth_id")
		if authIdStr == "" {
			return c.JSON(http.StatusBadRequest, NewErrorResponse("auth_id is required"))
		}

		user, err := authService.AuthenticateUser(model.AuthId(authIdStr))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, NewErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, user)
	}
}

func GetUserByID(authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userIdStr := c.QueryParam("user_id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, NewErrorResponse("user_id is required"))
		}

		user, err := authService.GetUserByID(model.UserId(userId))
		if err != nil {
			return c.JSON(http.StatusNotFound, NewErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, user)
	}
}
