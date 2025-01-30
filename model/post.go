package model

import "time"

type PostId int
type PrefectureId int

type Post struct {
	PostId    PostId
	PrefectureId PrefectureId
	Image        string
	Date         time.Time
	Comment      string
	Longitude    float64
	Latitude     float64
	UserId       UserId
	Groups       []GroupId
}

type UploadPost struct {
	Image        string
	Date         time.Time
	Comment      string
	PrefectureId PrefectureId
	Longitude    float64
	Latitude     float64
	UserId       UserId
	Groups       []GroupId
}

type SelectPost struct {
	PrefectureId PrefectureId
	Groups       []GroupId
}
