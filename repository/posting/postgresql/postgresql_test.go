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
			createdUser, createdPost := createUserAndPostForTest(t, userRepository, postRepository, tc.createUser, tc.uploadPost)

			tc.want.UserId = createdUser.UserId
			tc.want.PostId = createdPost.PostId

			if diff := cmp.Diff(tc.want, *createdPost); diff != "" {
				t.Errorf("CreatePost() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCreatePostNotUser(t *testing.T) {
	testCases := []struct {
		name       string
		uploadPost model.UploadPost
	}{
		{
			name: "ユーザーが存在しないUserIdでの画像の投稿",
			uploadPost: model.UploadPost{
				UserId:       111111111,
				Image:        "hoge",
				Date:         time.Date(2025, time.January, 30, 15, 4, 5, 0, time.UTC),
				Comment:      "hogehoge",
				PrefectureId: 1,
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

	postRepository := postPostgresql.NewPostRepository(db)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			gotPost, err := postRepository.CreatePost(tc.uploadPost)
			if err == nil {
				t.Errorf("CreatePost() expected an error but got none")
			}

			if gotPost != nil {
				t.Errorf("CreatePost() expected nil post but got %+v", gotPost)
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
			createdUser, createdPost := createUserAndPostForTest(t, userRepository, postRepository, tc.createUser, tc.uploadPost)

			tc.want.UserId = createdUser.UserId
			tc.want.PostId = createdPost.PostId

			if diff := cmp.Diff(tc.want, *createdPost); diff != "" {
				t.Errorf("CreatePost() mismatch (-want +got):\n%s", diff)
			}

			deletePost := model.DeletePost{
				PostId: createdPost.PostId,
				UserId: createdUser.UserId,
			}

			gotPostId, err := postRepository.DeletePost(deletePost)
			if err != nil {
				t.Fatalf("DeletePost() error = %v", err)
			}

			if gotPostId == nil || *gotPostId != deletePost.PostId {
				t.Errorf("DeletePost() PostId mismatch: expected %v, got %v", deletePost.PostId, gotPostId)
			}

			deletedPost, err := postRepository.FindPostByID(deletePost.PostId)
			if err == nil {
				t.Errorf("FindPostByID() expected error but got none")
			}
			if deletedPost != nil {
				t.Errorf("FindPostByID() expected nil but got %+v", deletedPost)
			}

			t.Logf("Test Passed: %s | Deleted PostID = %v", tc.name, deletePost.PostId)
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
			createdUser, createdPost := createUserAndPostForTest(t, userRepository, postRepository, tc.createUser, tc.uploadPost)

			tc.want.UserId = createdUser.UserId
			tc.want.PostId = createdPost.PostId

			gotPost, err := postRepository.FindPostByID(createdPost.PostId)
			if err != nil {
				t.Fatalf("FindPostByID() error = %v", err)
			}

			if diff := cmp.Diff(tc.want, *gotPost); diff != "" {
				t.Errorf("FetchPost() mismatch (-want +got):\n%s", diff)
			}

			t.Logf("Test Passed: %s | Created PostID = %v, Fetched PostID = %v",
				tc.name, createdPost.PostId, gotPost.PostId)
		})
	}
}

func TestFetchPostNotPost(t *testing.T) {
	testCases := []struct {
		name   string
		postId model.PostId
	}{
		{
			name:   "存在しないPostIdで投稿を取得しようとする",
			postId: 999999999,
		},
	}

	db, err := setupDB.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	postRepository := postPostgresql.NewPostRepository(db)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotPost, err := postRepository.FindPostByID(tc.postId)

			if err == nil {
				t.Errorf("FindPostByID() expected an error but got none")
			}

			if gotPost != nil {
				t.Errorf("FindPostByID() expected nil but got %+v", gotPost)
			}
		})
	}
}
