package handler

import (
	"fmt"
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"

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

			if _, err := groupService.RegisterPublicSetting(publicSetting); err != nil {
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
		prefectureName := c.QueryParam("prefecture_name")
		prefectureId, exists := model.PrefectureNameToId[prefectureName]
		if !exists {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid prefecture name"})
		}

		groupIdsStr := c.QueryParams()["groups"]
		var groupIds []model.GroupId
		for _, idStr := range groupIdsStr {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid group ID"})
			}
			groupIds = append(groupIds, model.GroupId(id))
		}

		prefectureIDAndGroupIDs := model.PrefectureIDAndGroupIDs{
			PrefectureId: prefectureId,
			GroupIds:     groupIds,
		}

		posts, err := postingService.GetPostsByPrefectureIDAndGroupIDs(prefectureIDAndGroupIDs)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		response := model.GetPostsResponse{}
		for _, post := range posts {
			response.Posts = append(response.Posts, model.PostResponse{
				PostId:    int(post.PostId),
				Image:     post.Image,
				Longitude: post.Longitude,
				Latitude:  post.Latitude,
			})
		}

		return c.JSON(http.StatusOK, response)
	}
}

func GetPost(postingService service.PostingService, authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		postIdStr := c.QueryParam("post_id")
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
		}

		post, err := postingService.GetPostByID(model.PostId(postId))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
		}

		user, err := authService.GetUserByID(post.UserId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user info"})
		}

		response := model.GetPostResponse{
			UserName: user.UserName,
			UserId:   post.UserId,
			Image:    post.Image,
			Comment:  post.Comment,
			Date:     post.Date,
		}

		return c.JSON(http.StatusOK, response)
	}
}
