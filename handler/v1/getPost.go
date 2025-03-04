package v1

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type getPostResponse struct {
	UserName model.UserName  `json:"user_name"`
	UserId   model.UserId    `json:"user_id"`
	Image    string    `json:"image"`
	Comment  string    `json:"comment"`
	Date     time.Time `json:"date"`
}

func GetPost(postingService service.PostingService, authService service.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		postIdStr := c.QueryParam("post_id")
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		}

		post, err := postingService.GetPostByID(model.PostId(postId))
		if err != nil {
			return c.JSON(http.StatusNotFound, model.NewErrorResponse(err.Error()))
		}

		user, err := authService.GetUserByID(post.UserId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		}

		response := getPostResponse{
			UserName: user.UserName,
			UserId:   post.UserId,
			Image:    post.Image,
			Comment:  post.Comment,
			Date:     post.Date,
		}

		return c.JSON(http.StatusOK, response)
	}
}
