package handler

import (
	"fmt"
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

func RegisterGroup(groupService service.GroupService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.CreateGroup
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, NewErrorResponse("Invalid request body"))
		}

		createdGroup, err := groupService.RegisterGroup(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		}

		joinGroup := model.JoinGroup{
			UserId:   req.AuthorId,
			GroupId:  createdGroup.GroupId,
			Password: req.Password,
		}

		joinedGroupId, err := groupService.JoinGroup(joinGroup)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, joinedGroupId)
	}
}

func JoinGroup(groupService service.GroupService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.JoinGroup
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, NewErrorResponse("Invalid request body"))
		}

		joinedGroupId, err := groupService.JoinGroup(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"group_id": joinedGroupId,
		})
	}
}

func GetUserGroup(groupService service.GroupService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userIdStr := c.QueryParam("user_id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, NewErrorResponse("Invalid user_id"))
		}

		userGroups, err := groupService.GetUserGroupSummaryByUserID(model.UserId(userId))
		if err != nil {
			return c.JSON(http.StatusNotFound, NewErrorResponse("User not found"))
		}

		response := model.GetUserGroupResponse{
			Groups: userGroups,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func GetGroupMembers(groupService service.GroupService, authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		groupIdStr := c.QueryParam("group_id")
		groupId, err := strconv.Atoi(groupIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, NewErrorResponse("Invalid group_id"))
		}

		groupMembers, err := groupService.GetGroupMembersByGroupID(model.GroupId(groupId))
		if err != nil {
			return c.JSON(http.StatusNotFound, NewErrorResponse("Group not found"))
		}

		response := model.GetGroupMembersResponse{}

		for _, userId := range groupMembers.Members {
			user, err := authService.GetUserByID(userId)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, NewErrorResponse(fmt.Sprintf("Failed to get user: %v", err)))
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

func GetGroup(groupService service.GroupService) echo.HandlerFunc {
	return func(c echo.Context) error {
		groupIdStr := c.QueryParam("group_id")
		groupId, err := strconv.Atoi(groupIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, NewErrorResponse("Invalid group_id"))
		}

		groupSummary, err := groupService.GetGroupSummaryByGroupID(model.GroupId(groupId))
		if err != nil {
			return c.JSON(http.StatusNotFound, NewErrorResponse("Group not found"))
		}

		response := model.GetGroupResponse{
			Name: groupSummary.GroupName,
			Icon: groupSummary.Icon,
		}

		return c.JSON(http.StatusOK, response)
	}
}
