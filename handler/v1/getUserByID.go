package v1

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type getUserByIDResponse struct {
	Name model.UserName`json:"name"`
	Icon string `json:"icon"`
}

func GetUserByID(authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userIdStr := c.QueryParam("user_id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		}

		user, err := authService.GetUserByID(model.UserId(userId))
		if err != nil {
			return c.JSON(http.StatusNotFound, model.NewErrorResponse(err.Error()))
		}

		response := getUserByIDResponse{
			Name: user.UserName,
			Icon: user.Icon,
		}
		return c.JSON(http.StatusOK, response)
	}
}
