package model

import "time"

type PostId int
type PrefectureId int

type Post struct {
	PostId       PostId
	PrefectureId PrefectureId
	Image        string
	Date         time.Time
	Comment      string
	Longitude    float64
	Latitude     float64
	UserId       UserId
}

type UploadPost struct {
	Image        string
	Date         time.Time
	Comment      string
	PrefectureId PrefectureId
	Longitude    float64
	Latitude     float64
	UserId       UserId
}

type DeletePost struct {
	PostId PostId
	UserId UserId
}

type PrefectureIDAndGroupID struct {
	PrefectureId PrefectureId
	GroupId      GroupId
}

type PrefectureIDAndGroupIDs struct {
	PrefectureId PrefectureId
	GroupIds     []GroupId
}

type UploadPostRequest struct {
	Image        string    `json:"image"`
	Date         time.Time `json:"date"`
	Comment      string    `json:"comment"`
	PrefectureId int       `json:"prefecture_id"`
	Longitude    float64   `json:"longitude"`
	Latitude     float64   `json:"latitude"`
	UserId       int       `json:"user_id"`
	Groups       []int     `json:"groups"`
}

type GetPostsRequest struct {
	PrefectureId string `json:"prefecture_id"`
	Groups       []int  `json:"groups"`
}

type GetPostsResponse struct {
	UserName UserName  `json:"user_name"`
	UserId   UserId      `json:"user_id"`
	Image    string    `json:"image"`
	Comment  string    `json:"comment"`
	Date     time.Time `json:"date"`
}
