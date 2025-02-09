package handler

import (
	"net/http"
	"painteer/model"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

func RegisterGroup(groupService service.GroupService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.CreateGroup
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		createdGroup, err := groupService.RegisterGroup(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		joinGroup := model.JoinGroup{
			UserId:   req.AuthorId,
			GroupId:  createdGroup.GroupId,
			Password: req.Password,
		}

		joinedGroupId, err := groupService.JoinGroup(joinGroup)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"group_id": joinedGroupId,
		})
	}
}
