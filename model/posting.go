package model

import "time"

type PostingId int
type PrefectureId int

type Posting struct {
	PostingId PostingId
	Image     string
	Date      time.Time
	Comment   string
	Region    string
	Longitude float64
	Latitude  float64
	UserID    UserId
}

type CreatePosting struct {
	Image     string
	Date      time.Time
	Comment   string
	Region    string
	Longitude float64
	Latitude  float64
	UserID    UserId
	Groups    []GroupId
}

type SelectPosting struct {
	PrefectureId PrefectureId
	Groups       []GroupId
}
