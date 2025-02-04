package postgresql_test

import (
	"painteer/model"
	userPostgresql "painteer/repository/auth/postgresql"
	groupPostgresql "painteer/repository/group/postgresql"
	setupDB "painteer/repository/utils"
	"testing"

	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq"
)

func createUserAndGroupForTest(
	t *testing.T,
	userRepository *userPostgresql.AuthRepositoryImpl,
	groupRepository *groupPostgresql.GroupRepositoryImpl,
	createUser model.CreateUser,
	createGroup model.CreateGroup,
) (*model.User, *model.Group, error) {
	t.Helper()

	createdUser, err := userRepository.CreateUser(createUser)
	if err != nil {
		t.Fatalf("CreateUser() error = %v", err)
	}
	if createdUser == nil {
		t.Fatal("CreateUser() returned nil, expected valid User")
	}

	createGroup.AuthorId = createdUser.UserId

	createdGroup, err := groupRepository.CreateGroup(createGroup)
	if err != nil {
		return nil, nil, err
	}

	return createdUser, createdGroup, nil
}

func TestCreateGroup(t *testing.T) {
	testCases := []struct {
		name         string
		createAuthor *model.CreateUser
		createGroup  model.CreateGroup
		want         *model.Group
		expectErr    bool
	}{
		{
			name: "ユーザー&グループの作成",
			createAuthor: &model.CreateUser{
				UserName: "GroupAuthor",
				Icon:     "hoge",
				AuthId:   "hoge",
			},
			createGroup: model.CreateGroup{
				GroupName: "testCreateGroup",
				Icon:      "hoge",
				Password:  "test",
			},
			want: &model.Group{
				GroupName: "testCreateGroup",
				Icon:      "hoge",
				Password:  "test",
			},
			expectErr: false,
		},
		{
			name:         "存在しないuserIdでのグループ作成",
			createAuthor: nil,
			createGroup: model.CreateGroup{
				GroupName: "testCreateGroup",
				Icon:      "hoge",
				Password:  "test",
				AuthorId:  123456789,
			},
			want:      nil,
			expectErr: true,
		},
	}

	db, err := setupDB.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	userRepository := userPostgresql.NewAuthRepository(db)
	groupRepository := groupPostgresql.NewGroupRepository(db)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var createdUser *model.User
			var gotGroup *model.Group
			var err error

			if tc.createAuthor != nil {
				createdUser, gotGroup, err = createUserAndGroupForTest(t, userRepository, groupRepository, *tc.createAuthor, tc.createGroup)
			} else {
				gotGroup, err = groupRepository.CreateGroup(tc.createGroup)
			}

			if tc.expectErr {
				if err == nil {
					t.Fatalf("Expected error, but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("CreateGroup() error = %v", err)
			}
			tc.want.AuthorId = createdUser.UserId
			tc.want.GroupId = gotGroup.GroupId

			if diff := cmp.Diff(tc.want, gotGroup); diff != "" {
				t.Fatalf("CreatedGroup() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
