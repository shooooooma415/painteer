package v1

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetGroup(groupService service.GroupService) echo.HandlerFunc {
	return func(c echo.Context) error {
		groupIdStr := c.QueryParam("group_id")
		groupId, err := strconv.Atoi(groupIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		}

		groupSummary, err := groupService.GetGroupSummaryByGroupID(model.GroupId(groupId))
		if err != nil {
			return c.JSON(http.StatusNotFound, model.NewErrorResponse(err.Error()))
		}

		response := model.GetGroupResponse{
			Name: groupSummary.GroupName,
			Icon: groupSummary.Icon,
		}

		return c.JSON(http.StatusOK, response)
	}
}
