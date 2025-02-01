package postgresql_test

import (
	"painteer/model"
	postgresql "painteer/repository/auth/postgresql"
	setupDB "painteer/repository/utils"
	"testing"

	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq"
)

func TestCreateAndFindUserByUserID(t *testing.T) {
	testCases := []struct {
		name       string
		createUser model.CreateUser
		want       model.User
	}{
		{
			name: "ユーザーの作成&検索",
			createUser: model.CreateUser{
				UserName: "hoge",
				Icon:     "hogehoge",
				AuthId:   "auth_hoge1",
			},
			want: model.User{
				UserName: "hoge",
				Icon:     "hogehoge",
				AuthId:   "auth_hoge1",
			},
		}}

	db, err := setupDB.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	repository := postgresql.NewAuthRepository(db)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			createdUser, err := repository.CreateUser(tc.createUser)
			if err != nil {
				t.Fatalf("CreateUser() error = %v", err)
			}
			if createdUser == nil {
				t.Fatal("CreateUser() returned nil, expected valid User")
			}

			tc.want.UserId = createdUser.UserId
			
			if diff := cmp.Diff(tc.want, *createdUser); diff != "" {
				t.Errorf("CreateUser() mismatch (-want +got):\n%s", diff)
			}

			gotUser, err := repository.FindUserByUserID(createdUser.UserId)
			if err != nil {
				t.Fatalf("FindUserByID() error = %v", err)
			}

			if diff := cmp.Diff(tc.want, *gotUser); diff != "" {
				t.Errorf("FindUserByUserID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCreateAndFindUserByAuthID(t *testing.T) {
	testCases := []struct {
		name       string
		createUser model.CreateUser
		want       model.User
	}{
		{
			name: "ユーザーの作成&検索",
			createUser: model.CreateUser{
				UserName: "hoge",
				Icon:     "hogehoge",
				AuthId:   "auth_hoge2",
			},
			want: model.User{
				UserName: "hoge",
				Icon:     "hogehoge",
				AuthId:   "auth_hoge2",
			},
		}}

	db, err := setupDB.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	repository := postgresql.NewAuthRepository(db)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			createdUser, err := repository.CreateUser(tc.createUser)
			if err != nil {
				t.Fatalf("CreateUser() error = %v", err)
			}
			if createdUser == nil {
				t.Fatal("CreateUser() returned nil, expected valid User")
			}

			tc.want.UserId = createdUser.UserId

			if diff := cmp.Diff(tc.want, *createdUser); diff != "" {
				t.Errorf("CreateUser() mismatch (-want +got):\n%s", diff)
			}

			gotUser, err := repository.FindUserByAuthID(tc.createUser.AuthId)
			if err != nil {
				t.Fatalf("FindUserByID() error = %v", err)
			}

			if diff := cmp.Diff(tc.want, *gotUser); diff != "" {
				t.Errorf("FindUserByAuthID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
