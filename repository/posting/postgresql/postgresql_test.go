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
			tc.uploadPost.UserId = createdUser.UserId
			
			gotPost, err := postRepository.CreatePost(tc.uploadPost)
			

			if err != nil {
				t.Fatalf("CreatedPost() error = %v", err)
			}
			
			tc.want.PostId = gotPost.PostId

			if diff := cmp.Diff(tc.want, *gotPost); diff != "" {
				t.Errorf("CreatePost() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCreateUserAndPostAndDeletePost(t *testing.T) {
	testCases := []struct {
		name       string
		createUser model.CreateUser
		uploadPost model.UploadPost
		want       model.Post
	}{
		{
			name: "ユーザーの作成＆画像の投稿&投稿の削除",
			createUser: model.CreateUser{
				UserName: "hoge",
				Icon:     "hoge",
				AuthId:   "hogehogehoge1",
			},
			uploadPost: model.UploadPost{
				Image:        "hoge",
				Date:         time.Date(2025, time.January, 30, 15, 4, 5, 0, time.UTC),
				Comment:      "hogehoge",
				PrefectureId: 1,
				Longitude:    123.456,
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
			tc.uploadPost.UserId = createdUser.UserId
			gotPost, err := postRepository.CreatePost(tc.uploadPost)

			if err != nil {
				t.Fatalf("CreatePost() error = %v", err)
			}
			
			deletePost:= model.DeletePost{
				PostId: gotPost.PostId,
				UserId: createdUser.UserId,
			}

			gotPostId, err := postRepository.DeletePost(deletePost)
			tc.want.PostId = deletePost.PostId

			if err != nil {
				t.Fatalf("DeletePost() error = %v", err)
			}
			if diff := cmp.Diff(tc.want.PostId, *gotPostId); diff != "" {
				t.Errorf("DeletePost() mismatch (-want +got):\n%s", diff)
			}

			deletedPost, err := postRepository.FindPostByID(deletePost.PostId)
			if err == nil {
				t.Errorf("FindPostByID() expected error but got none")
			}
			if deletedPost != nil {
				t.Errorf("FindPostByID() expected nil but got %+v", deletedPost)
			}

		})
	}
}

func TestCreateUserAndPostAndFetchPost(t *testing.T) {
	testCases := []struct {
		name       string
		createUser model.CreateUser
		uploadPost model.UploadPost
		want       model.Post
	}{
		{
			name: "ユーザーの作成＆画像の投稿&投稿のfetch",
			createUser: model.CreateUser{
				UserName: "hoge",
				Icon:     "hoge",
				AuthId:   "hogehogehoge2",
			},
			uploadPost: model.UploadPost{
				Image:        "hoge",
				Date:         time.Date(2025, time.January, 30, 15, 4, 5, 0, time.UTC),
				Comment:      "hogehoge",
				PrefectureId: 1,
				Longitude:    123.456,
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
			tc.uploadPost.UserId = createdUser.UserId
			createdPost, err := postRepository.CreatePost(tc.uploadPost)

			if err != nil {
				t.Fatalf("CreatePost() error = %v", err)
			}

			tc.want.PostId = createdPost.PostId

			if diff := cmp.Diff(tc.want, *createdPost); diff != "" {
				t.Errorf("CreatePost() mismatch (-want +got):\n%s", diff)
			}

			gotPost,err := postRepository.FindPostByID(createdPost.PostId)

			if err != nil {
				t.Fatalf("FetchPost() error = %v", err)
			}

			if diff := cmp.Diff(tc.want, *gotPost); diff != "" {
				t.Errorf("FetchPost() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
