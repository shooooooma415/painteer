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

func GetPosts(postingService service.PostingService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.GetPostsRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		prefectureId, exists := model.PrefectureNameToId[req.PrefectureId]
		if !exists {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid prefecture name"})
		}
		groupIds := make([]model.GroupId, len(req.Groups))
		for i, id := range req.Groups {
			groupIds[i] = model.GroupId(id)
		}

		prefectureIDAndGroupIDs := model.PrefectureIDAndGroupIDs{
			PrefectureId: prefectureId,
			GroupIds:     groupIds,
		}

		posts, err := postingService.GetPostsByPrefectureIDAndGroupIDs(prefectureIDAndGroupIDs)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, posts)
	}
}

func GetPost(postingService service.PostingService, authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.PostId
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		post, err := postingService.GetPostByID(model.PostId(req))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
		}

		user, err := authService.GetUserByID(post.UserId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user info"})
		}

		response := model.GetPostsResponse{
			UserName: user.UserName,
			UserId:   post.UserId,
			Image:    post.Image,
			Comment:  post.Comment,
			Date:     post.Date,
		}

		return c.JSON(http.StatusOK, response)
	}
}
