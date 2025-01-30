package model

import "time"

type PostId int
type PrefectureId int

type Post struct {
	PostingId PostId
	Image     string
	Date      time.Time
	Comment   string
	Region    string
	Longitude float64
	Latitude  float64
	UserID    UserId
}

type CreatePost struct {
	Image     string
	Date      time.Time
	Comment   string
	Region    string
	Longitude float64
	Latitude  float64
	UserID    UserId
	Groups    []GroupId
}

type SelectPost struct {
	PrefectureId PrefectureId
	Groups       []GroupId
}
