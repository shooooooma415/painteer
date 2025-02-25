package v1

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type getUserGroupResponse struct {
	Groups []model.GroupSummary `json:"groups"`
}



func GetUserGroup(groupService service.GroupService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userIdStr := c.QueryParam("user_id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		}

		userGroups, err := groupService.GetUserGroupSummaryByUserID(model.UserId(userId))
		if err != nil {
			return c.JSON(http.StatusNotFound, model.NewErrorResponse(err.Error()))
		}

		response := getUserGroupResponse{
			Groups: userGroups,
		}

		return c.JSON(http.StatusOK, response)
	}
}
