package model

type CountPostByPrefectureID struct {
	PrefectureId int
	PostCount    int
}

type Count struct {
	Data []CountPostByPrefectureID
}
