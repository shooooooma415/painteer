package postgresql_test

import (
	"painteer/model"
	userPostgresql "painteer/repository/auth/postgresql"
	postPostgresql "painteer/repository/posting/postgresql"
	setupDB "painteer/repository/utils"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq"
)

func TestCreateUserAndPost(t *testing.T) {
	testCases := []struct {
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

	db, err := setupDB.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	userRepository := userPostgresql.NewAuthRepository(db)
	postRepository := postPostgresql.NewPostRepository(db)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			createdUser, err := userRepository.CreateUser(tc.createUser)

			if err != nil {
				t.Fatalf("CreateUser() error = %v", err)
			}
			if createdUser == nil {
				t.Fatal("CreateUser() returned nil, expected valid User")
			}

			tc.want.UserId = createdUser.UserId
			gotPost, err := postRepository.CreatePost(tc.uploadPost)

			if err != nil {
				t.Fatalf("FindUserByID() error = %v", err)
			}

			if diff := cmp.Diff(tc.want, *gotPost); diff != "" {
				t.Errorf("CreatePost() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCreateUserAndPostAndDeletePost(t *testing.T) {

}

func TestCreateUserAndPostAndFetchPost(t *testing.T) {

}
