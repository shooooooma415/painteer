package v1

import (
	"net/http"
	"painteer/model"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

func JoinGroup(groupService service.GroupService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.JoinGroup
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		}

		joinedGroupId, err := groupService.JoinGroup(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"group_id": joinedGroupId,
		})
	}
}
