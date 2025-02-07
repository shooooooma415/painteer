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


type PrefectureIDAndGroupID struct{
	PrefectureId PrefectureId
	GroupId GroupId
}

type PrefectureIDAndGroupIDs struct {
	PrefectureId PrefectureId
	GroupIds     []GroupId
}

type Posts struct {
	Posts []Post
}