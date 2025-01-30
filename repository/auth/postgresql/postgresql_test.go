package postgresql_test

import (
	"painteer/model"
	postgresql "painteer/repository/auth/postgresql"
	setupDB "painteer/repository/utils"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name       string
		createUser model.CreateUser
		want       model.User
	}{
		{
			name: "正常にユーザーを作成する",
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

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			db, err := setupDB.ConnectDB()
			if err != nil {
				t.Fatalf("Failed to connect to the database: %v", err)
			}
			defer db.Close()

			repository := postgresql.NewAuthRepository(db)
			got, err := repository.CreateUser(tc.createUser)

			if err != nil {
				t.Fatalf("CreateUserQuery() error = %v", err)
			}

			if got.UserName != tc.want.UserName {
				t.Errorf("CreateUser() UserName = %v, want %v", got.UserName, tc.want.UserName)
			}
			if got.AuthId != tc.want.AuthId {
				t.Errorf("CreateUser() AuthId = %v, want %v", got.AuthId, tc.want.AuthId)
			}
			if got.Icon != tc.want.Icon {
				t.Errorf("CreateUser() Icon = %v, want %v", got.Icon, tc.want.Icon)
			}

		})
	}
}

func TestFindUserByID(t *testing.T) {
	testCases := []struct {
		name   string
		authId model.AuthId
	}{
		name:   "authIdを元に正常にuserIdを取得する",
		authId: "hoge",
	}
}
