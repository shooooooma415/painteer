package model

type CountByPrefectureID struct {
	PrefectureId int
	PostCount    int
}

type Count struct {
	Data []CountByPrefectureID
}
