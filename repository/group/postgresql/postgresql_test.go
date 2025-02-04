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

func TestCreateAndJoinGroup(t *testing.T) {
	testCases := []struct {
		name         string
		createAuthor *model.CreateUser
		createGroup  *model.CreateGroup
		joinGroup    *model.JoinGroup
		want         *model.JoinGroup
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
			want:      &model.JoinGroup{},
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
			joinGroup: &model.JoinGroup{
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
			joinGroup: &model.JoinGroup{
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

			if tc.joinGroup == nil {
				tc.joinGroup = &model.JoinGroup{
					UserId:  createdUser.UserId,
					GroupId: createdGroup.GroupId,
				}
			}

			gotJoinGroup, err := groupRepository.JoinGroup(*tc.joinGroup)

			if tc.expectErr {
				if err == nil {
					t.Fatalf("Expected error, but got nil")
				} else {
					t.Logf("Correctly returned error: %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("JoinGroup() error = %v", err)
			}

			tc.want.UserId = tc.joinGroup.UserId
			tc.want.GroupId = tc.joinGroup.GroupId

			if diff := cmp.Diff(tc.want, gotJoinGroup); diff != "" {
				t.Fatalf("JoinGroup() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCreateUserAndGroupAndFindGroupID(t *testing.T) {
	testCases := []struct {
		name           string
		createAuthor   model.CreateUser
		createGroup    model.CreateGroup
		verifyPassword model.VerifyPassword
		want           *model.Group
		expectErr      bool
	}{
		{
			name: "ユーザー&グループの作成->passwordとgroupNameが一致している",
			createAuthor: model.CreateUser{
				UserName: "GroupAuthor",
				Icon:     "hoge",
				AuthId:   "hoge",
			},
			createGroup: model.CreateGroup{
				GroupName: "testCreateGroup1",
				Icon:      "hoge",
				Password:  "test",
			},
			verifyPassword: model.VerifyPassword{
				Password:  "test",
				GroupName: "testCreateGroup1",
			},
			want: &model.Group{
				GroupName: "testCreateGroup",
				Icon:      "hoge",
				Password:  "test",
			},
			expectErr: false,
		},
		{
			name: "ユーザー&グループの作成->passwordが間違っている",
			createAuthor: model.CreateUser{
				UserName: "GroupAuthor",
				Icon:     "hoge",
				AuthId:   "hoge",
			},
			createGroup: model.CreateGroup{
				GroupName: "testCreateGroup2",
				Icon:      "hoge",
				Password:  "test",
			},
			verifyPassword: model.VerifyPassword{
				Password:  "falsePassword",
				GroupName: "testCreateGroup2",
			},
			want:      nil,
			expectErr: true,
		},
		{
			name: "ユーザー&グループの作成->groupNameが間違っている",
			createAuthor: model.CreateUser{
				UserName: "GroupAuthor",
				Icon:     "hoge",
				AuthId:   "hoge",
			},
			createGroup: model.CreateGroup{
				GroupName: "testCreateGroup3",
				Icon:      "hoge",
				Password:  "test",
			},
			verifyPassword: model.VerifyPassword{
				Password:  "test",
				GroupName: "falseTestCreateGroup3",
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
			var createdGroup *model.Group
			var err error

			if tc.createGroup.GroupName != "" {
				_, createdGroup, err = createUserAndGroupForTest(t, userRepository, groupRepository, tc.createAuthor, tc.createGroup)
				if err != nil {
					t.Fatalf("Failed to create user and group: %v", err)
				}
			}

			groupId, err := groupRepository.FindGroupIDByPasswordAndName(tc.verifyPassword)
			if tc.expectErr {
				if err == nil {
					t.Fatalf("Expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("FindGroupIDByPasswordAndName() error = %v", err)
			}

			if createdGroup != nil {
				tc.want.GroupId = createdGroup.GroupId
			}
			if diff := cmp.Diff(tc.want.GroupId, *groupId); diff != "" {
				t.Fatalf("FindGroupIDByPasswordAndName() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}