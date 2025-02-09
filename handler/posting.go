package handler

import (
	"fmt"
	"net/http"
	"painteer/model"
	"painteer/service"

	"github.com/labstack/echo/v4"
)

func UploadPost(postingService service.PostingService, groupService service.GroupService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.UploadPostRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		uploadPost := model.UploadPost{
			Image:        req.Image,
			Date:         req.Date,
			Comment:      req.Comment,
			PrefectureId: model.PrefectureId(req.PrefectureId),
			Longitude:    req.Longitude,
			Latitude:     req.Latitude,
			UserId:       model.UserId(req.UserId),
		}

		createdPost, err := postingService.CreatePost(uploadPost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		for _, groupId := range req.Groups {
			publicSetting := model.PublicSetting{
				PostId:        *createdPost,
				PublicGroupId: model.GroupId(groupId),
			}

			if err := groupService.PublicSetting(publicSetting); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": fmt.Sprintf("Failed to set public setting for group %d", groupId),
				})
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"is_success": true,
		})
	}
}