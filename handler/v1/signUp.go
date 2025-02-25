package v1

import (
	"net/http"
	"painteer/model"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

type signUpResponse struct {
	UserId model.UserId `json:"user_id"`
}

func SignUp(authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.CreateUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		}

		user, err := authService.RegisterUser(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		}

		response := signUpResponse{
			UserId: user.UserId,
		}

		return c.JSON(http.StatusOK, response)
	}
}
