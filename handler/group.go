package handler

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"

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

		return c.JSON(http.StatusOK, joinedGroupId)
	}
}

func JoinGroup(groupService service.GroupService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.JoinGroup
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		joinedGroupId, err := groupService.JoinGroup(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"group_id": joinedGroupId,
		})
	}
}

func GetGroupMembers(groupService service.GroupService, authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		groupIdStr := c.QueryParam("group_id")
		groupId, err := strconv.Atoi(groupIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid group_id"})
		}

		groupMembers, err := groupService.GetGroupMembersByGroupID(model.GroupId(groupId))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Group not found"})
		}

		response := model.GetGroupMembersResponse{}

		for _, userId := range groupMembers.Members {
			user, err := authService.GetUserByID(userId)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to fetch user info for user_id %d", userId)})
			}

			response.Member = append(response.Member, struct {
				UserId   model.UserId   `json:"user_id"`
				UserName model.UserName `json:"user_name"`
				Icon     string         `json:"icon"`
			}{
				UserId:   user.UserId,
				UserName: user.UserName,
				Icon:     user.Icon,
			})
		}

		return c.JSON(http.StatusOK, response)
	}
}
