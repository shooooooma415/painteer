package postgresql_test

import (
	"painteer/model"
	userPostgresql "painteer/repository/auth/postgresql"
	postPostgresql "painteer/repository/posting/postgresql"
	paintPostgresql "painteer/repository/paint/postgresql"
	setupDB "painteer/repository/utils"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq"
)

func createUserAndPostForTest(
	t *testing.T,
	userRepository *userPostgresql.AuthRepositoryImpl,
	postRepository *postPostgresql.PostRepositoryImpl,
	createUser model.CreateUser,
	uploadPost model.UploadPost,
) (*model.User, *model.Post) {
	t.Helper()
	createdUser, err := userRepository.CreateUser(createUser)
	if err != nil {
		t.Fatalf("CreateUser() error = %v", err)
	}
	if createdUser == nil {
		t.Fatal("CreateUser() returned nil, expected valid User")
	}
	t.Logf("Created User: %+v", createdUser)

	uploadPost.UserId = createdUser.UserId
	createdPost, err := postRepository.CreatePost(uploadPost)
	if err != nil {
		t.Fatalf("CreatePost() error = %v", err)
	}
	t.Logf("Created Post: %+v", createdPost)

	return createdUser, createdPost
}

func TestCreateUserAndPostAndFetchPostCount(t *testing.T) {
	testCases := []struct {
		name       string
		createUser model.CreateUser
		uploadPost model.UploadPost
		want       model.Count
	}{
		{
			name: "ユーザーの作成＆画像の投稿",
			createUser: model.CreateUser{
				UserName: "hoge",
				Icon:     "hoge",
				AuthId:   "hogehogehogehoge",
			},
			uploadPost: model.UploadPost{
				Image:        "hoge",
				Date:         time.Date(2025, time.January, 30, 15, 4, 5, 0, time.UTC),
				Comment:      "hogehoge",
				PrefectureId: 1,
				Longitude:    123.456,
				Latitude:     123.456,
			},
			want: model.Count{
				Data: []model.CountByPrefectureID{
					{
						PrefectureId: 1,
						PostCount:    1,
					},
				},
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
	pantRepository := paintPostgresql.NewPaintRepository(db)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// ユーザーと投稿の作成
			_, createdPost := createUserAndPostForTest(t, userRepository, postRepository, tc.createUser, tc.uploadPost)

			// 投稿数を取得
			count, err := pantRepository.CountPostsByPrefecture(createdPost.GroupId)
			if err != nil {
				t.Fatalf("CountPostsByPrefecture() error = %v", err)
			}

			// 結果を比較
			if diff := cmp.Diff(tc.want, *count); diff != "" {
				t.Errorf("CountPostsByPrefecture() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
