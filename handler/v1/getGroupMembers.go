package v1

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type getGroupMembersResponse struct {
	Member []struct {
		UserId   model.UserId   `json:"user_id"`
		UserName model.UserName `json:"user_name"`
		Icon     string   `json:"icon"`
	} `json:"member"`
}


func GetGroupMembers(groupService service.GroupService, authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		groupIdStr := c.QueryParam("group_id")
		groupId, err := strconv.Atoi(groupIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		}

		groupMembers, err := groupService.GetGroupMembersByGroupID(model.GroupId(groupId))
		if err != nil {
			return c.JSON(http.StatusNotFound, model.NewErrorResponse(err.Error()))
		}

		response := getGroupMembersResponse{}

		for _, userId := range groupMembers.Members {
			user, err := authService.GetUserByID(userId)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
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
