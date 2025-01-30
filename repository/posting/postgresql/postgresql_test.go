package postgresql_test

import (
	"painteer/model"
	postgresql "painteer/repository/post/postgresql"
	setupDB "painteer/repository/utils"
	"testing"

	_ "github.com/lib/pq"
)

func TestUploadAndDeletePost(t *testing.T){
	// testCases := []struct{
	// 	name string
	// 	uploadPost model.UploadPost
	// }{
	// 	name: "画像の投稿＆削除",
	// 	uploadPost: model.UploadPost{
	// 		Image:"hogehoge",
	// 		Date:"2025-01-30T15:04:05Z",
	// 		Comment:"hoge",
	// 		PrefectureId:1,
	// 		Longitude:135,
	// 		Latitude:35,
	// 		UserId:??,
	// 		Groups:,
	// 	}
	// }
}