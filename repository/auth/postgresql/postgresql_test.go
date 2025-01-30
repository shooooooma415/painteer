package postgresql_test

import (
	"painteer/model"
	postgresql "painteer/repository/auth/postgresql"
	setupDB "painteer/repository/utils"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateAndFindUser(t *testing.T) {
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
				AuthId:   "auth_hoge",
			},
			want: model.User{
				UserName: "hoge",
				Icon:     "hogehoge",
				AuthId:   "auth_hoge",
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

			if createdUser.UserName != tc.want.UserName {
				t.Errorf("CreateUser() UserName = %v, want %v", createdUser.UserName, tc.want.UserName)
			}
			if createdUser.AuthId != tc.want.AuthId {
				t.Errorf("CreateUser() AuthId = %v, want %v", createdUser.AuthId, tc.want.AuthId)
			}
			if createdUser.Icon != tc.want.Icon {
				t.Errorf("CreateUser() Icon = %v, want %v", createdUser.Icon, tc.want.Icon)
			}

			tc.want.UserId = createdUser.UserId

			foundUser, err := repository.FindUserByID(tc.createUser.AuthId)
			if err != nil {
				t.Fatalf("FindUserByID() error = %v", err)
			}

			if foundUser.UserId != tc.want.UserId {
				t.Errorf("FindUserByID() UserId = %v, want %v", foundUser.UserId, tc.want.UserId)
			}
			if foundUser.UserName != tc.want.UserName {
				t.Errorf("FindUserByID() UserName = %v, want %v", foundUser.UserName, tc.want.UserName)
			}
			if foundUser.AuthId != tc.want.AuthId {
				t.Errorf("FindUserByID() AuthId = %v, want %v", foundUser.AuthId, tc.want.AuthId)
			}
			if foundUser.Icon != tc.want.Icon {
				t.Errorf("FindUserByID() Icon = %v, want %v", foundUser.Icon, tc.want.Icon)
			}
		})
	}
}
