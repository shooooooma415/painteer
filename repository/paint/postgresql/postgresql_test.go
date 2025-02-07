package postgresql_test

import (
	"painteer/model"
	userPostgresql "painteer/repository/auth/postgresql"
	postPostgresql "painteer/repository/posting/postgresql"
	paintPostgresql "painteer/repository/paint/postgresql"
	setupDB "painteer/repository/utils"
	testUtils "painteer/repository/TestUtils"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq"
)

func TestFindPostIDsByPrefecture(t *testing.T) {
	testCases := []struct {
		name       string
		createUser model.CreateUser
		uploadPosts []model.UploadPost
		want       []model.PostsByPrefecture
	}{
		{
			name: "ユーザーの作成＆複数の投稿->都道府県ごとの投稿IDリストを取得",
			createUser: model.CreateUser{
				UserName: "hoge",
				Icon:     "hoge",
				AuthId:   "hogehogehogehoge",
			},
			uploadPosts: []model.UploadPost{
				{
					Image:        "post1",
					Date:         time.Date(2025, time.January, 30, 15, 4, 5, 0, time.UTC),
					Comment:      "hogehoge1",
					PrefectureId: 1,
					Longitude:    123.456,
					Latitude:     123.456,
				},
				{
					Image:        "post2",
					Date:         time.Date(2025, time.January, 31, 10, 10, 10, 0, time.UTC),
					Comment:      "hogehoge2",
					PrefectureId: 1,
					Longitude:    124.456,
					Latitude:     124.456,
				},
				{
					Image:        "post3",
					Date:         time.Date(2025, time.February, 1, 12, 0, 0, 0, time.UTC),
					Comment:      "hogehoge3",
					PrefectureId: 2,
					Longitude:    125.456,
					Latitude:     125.456,
				},
			},
			want: []model.PostsByPrefecture{
				{
					PrefectureId: 1,
					PostIds:      []model.PostId{}, // 実際のPostIdはテスト中に取得
				},
				{
					PrefectureId: 2,
					PostIds:      []model.PostId{},
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
	paintRepository := paintPostgresql.NewPaintRepository(db)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var createdPostIDs []model.PostId
			for _, post := range tc.uploadPosts {
				_, createdPost := testUtils.CreateUserAndPostForTest(t, userRepository, postRepository, tc.createUser, post)
				createdPostIDs = append(createdPostIDs, createdPost.PostId)
			}

			expectedResult := make(map[model.PrefectureId][]model.PostId)
			for i, post := range tc.uploadPosts {
				expectedResult[post.PrefectureId] = append(expectedResult[post.PrefectureId], createdPostIDs[i])
			}

			groupId := model.GroupId(1)
			result, err := paintRepository.FindPostIDsByPrefecture(groupId)
			if err != nil {
				t.Fatalf("FindPostIDsByPrefecture() error = %v", err)
			}

			for i, expected := range tc.want {
				if postIds, ok := expectedResult[expected.PrefectureId]; ok {
					tc.want[i].PostIds = postIds
				}
			}

			if diff := cmp.Diff(tc.want, result); diff != "" {
				t.Errorf("FindPostIDsByPrefecture() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}