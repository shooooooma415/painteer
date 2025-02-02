package postgresql_test

import (
	"painteer/model"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestCreateUserAndPost(t *testing.T) {
	testCase := []struct {
		name       string
		createUser model.CreateUser
		uploadPost model.UploadPost
		want       model.Post
	}{
		{
			name: "ユーザーの作成＆画像の投稿",
			createUser: model.CreateUser{
				UserName: "hoge",
				Icon:     "hoge",
				AuthId:   "hogehogehoge",
			},
			uploadPost: model.UploadPost{
				Image:        "hoge",
				Date:         time.Date(2025, time.January, 30, 15, 4, 5, 0, time.UTC),
				Comment:      "hogehoge",
				PrefectureId: 1,
				Longitude:    124.456,
				Latitude:     123.456,
			},
			want: model.Post{
				PrefectureId: 1,
				Image:        "hoge",
				Date:         time.Date(2025, time.January, 30, 15, 4, 5, 0, time.UTC),
				Comment:      "hogehoge",
				Longitude:    123.456,
				Latitude:     123.456,
			},
		},
	}
}

func TestCreateUserAndPostAndDeletePost(t *testing.T) {

}

func TestCreateUserAndPostAndFetchPost(t *testing.T) {

}
