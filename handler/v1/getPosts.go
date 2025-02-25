package v1

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPosts(postingService service.PostingService) echo.HandlerFunc {
	return func(c echo.Context) error {
		prefectureName := c.QueryParam("prefecture_name")
		prefectureId, exists := model.PrefectureNameToId[prefectureName]
		if !exists {
			return c.JSON(http.StatusBadRequest, model.NewErrorResponse("prefecture_name is invalid"))
		}

		groupIdsStr := c.QueryParams()["groups"]
		var groupIds []model.GroupId
		for _, idStr := range groupIdsStr {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
			}
			groupIds = append(groupIds, model.GroupId(id))
		}

		prefectureIDAndGroupIDs := model.PrefectureIDAndGroupIDs{
			PrefectureId: prefectureId,
			GroupIds:     groupIds,
		}

		posts, err := postingService.GetPostsByPrefectureIDAndGroupIDs(prefectureIDAndGroupIDs)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
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
