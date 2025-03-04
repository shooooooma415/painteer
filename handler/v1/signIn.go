package v1

import (
	"net/http"
	"painteer/model"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

type signInResponse struct {
	UserId model.UserId `json:"user_id"`
}


func SignIn(authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		authIdStr := c.QueryParam("auth_id")
		if authIdStr == "" {
			return c.JSON(http.StatusBadRequest, model.NewErrorResponse("auth_id is required"))
		}

		user, err := authService.AuthenticateUser(model.AuthId(authIdStr))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, model.NewErrorResponse(err.Error()))
		}

		response := signInResponse{
			UserId: user.UserId,
		}

		return c.JSON(http.StatusOK, response)
	}
}
