package v1

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"time"

	"github.com/labstack/echo/v4"
)

type uploadPostRequest struct {
	Image        string    `json:"image"`
	Date         time.Time `json:"date"`
	Comment      string    `json:"comment"`
	PrefectureId int       `json:"prefecture_id"`
	Longitude    float64   `json:"longitude"`
	Latitude     float64   `json:"latitude"`
	UserId       int       `json:"user_id"`
	Groups       []int     `json:"groups"`
}

func UploadPost(postingService service.PostingService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req uploadPostRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
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

		groupIds := make([]model.GroupId, len(req.Groups))
		for i, id := range req.Groups {
			groupIds[i] = model.GroupId(id)
		}

		_, err := postingService.CreatePost(uploadPost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		}

		response := model.UploadPostResponse{
			IsSuccess: true,
		}

		return c.JSON(http.StatusOK, response)
	}
}
