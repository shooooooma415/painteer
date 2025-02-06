package postgresql_test

import (
	"painteer/model"
	userPostgresql "painteer/repository/auth/postgresql"
	groupPostgresql "painteer/repository/group/postgresql"
	setupDB "painteer/repository/utils"
	testUtils "painteer/repository/testUtils"
	"testing"

	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq"
)


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
				createdUser, gotGroup, err = testUtils.CreateUserAndGroupForTest(t, userRepository, groupRepository, *tc.createAuthor, tc.createGroup)
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

func TestCreateAndInsertGroup(t *testing.T) {
	testCases := []struct {
		name         string
		createAuthor *model.CreateUser
		createGroup  *model.CreateGroup
		insertGroup    *model.InsertGroup
		want         *model.InsertGroup
		expectErr    bool
	}{
		{
			name: "ユーザー&グループの作成->作成者が参加",
			createAuthor: &model.CreateUser{
				UserName: "GroupAuthor",
				Icon:     "hoge",
				AuthId:   "hoge",
			},
			createGroup: &model.CreateGroup{
				GroupName: "testCreateGroup",
				Icon:      "hoge",
				Password:  "test",
			},
			want:      &model.InsertGroup{},
			expectErr: false,
		},
		{
			name: "存在しないグループに対して参加を試みる",
			createAuthor: &model.CreateUser{
				UserName: "GroupAuthor",
				Icon:     "hoge",
				AuthId:   "hoge",
			},
			createGroup: nil,
			insertGroup: &model.InsertGroup{
				GroupId: 12345678234567,
			},
			want:      nil,
			expectErr: true,
		},
		{
			name: "存在しない userId で参加を試みる",
			createAuthor: &model.CreateUser{
				UserName: "GroupAuthor",
				Icon:     "hoge",
				AuthId:   "hoge",
			},
			createGroup: &model.CreateGroup{
				GroupName: "testCreateGroup",
				Icon:      "hoge",
				Password:  "test",
			},
			InsertGroup: &model.InsertGroup{
				UserId: 12345678234567,
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
			var createdGroup *model.Group
			var err error

			if tc.createAuthor != nil {
				createdUser, err = userRepository.CreateUser(*tc.createAuthor)
				if err != nil {
					t.Fatalf("CreateUser() error = %v", err)
				}
			}

			if tc.createGroup != nil {
				if createdUser == nil {
					t.Fatalf("User must be created before creating a group")
				}
				tc.createGroup.AuthorId = createdUser.UserId
				createdGroup, err = groupRepository.CreateGroup(*tc.createGroup)
				if err != nil {
					t.Fatalf("CreateGroup() error = %v", err)
				}
			}

			if tc.insertGroup == nil {
				tc.insertGroup = &model.InsertGroup{
					UserId:  createdUser.UserId,
					GroupId: createdGroup.GroupId,
				}
			}

			gotInsertGroup, err := groupRepository.InsertGroup(*tc.insertGroup)

			if tc.expectErr {
				if err == nil {
					t.Fatalf("Expected error, but got nil")
				} else {
					t.Logf("Correctly returned error: %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("InsertGroup() error = %v", err)
			}

			tc.want.UserId = tc.insertGroup.UserId
			tc.want.GroupId = tc.insertGroup.GroupId

			if diff := cmp.Diff(tc.want, gotInsertGroup); diff != "" {
				t.Fatalf("InsertGroup() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
