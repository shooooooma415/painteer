package handler

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SignUp(authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.CreateUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		user, err := authService.RegisterUser(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		response := model.SignUpResponse{
			UserId: int(user.UserId),
		}

		return c.JSON(http.StatusOK, response)
	}
}

func SignIn(authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		authIdStr := c.QueryParam("auth_id")
		if authIdStr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing auth_id"})
		}

		user, err := authService.AuthenticateUser(model.AuthId(authIdStr))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		}

		return c.JSON(http.StatusOK, user)
	}
}

func GetUserByID(authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userIdStr := c.QueryParam("user_id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user_id"})
		}

		user, err := authService.GetUserByID(model.UserId(userId))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}

		return c.JSON(http.StatusOK, user)
	}
}
