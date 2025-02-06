package testUtils

import (
	"painteer/model"
	userPostgresql "painteer/repository/auth/postgresql"
	postPostgresql "painteer/repository/posting/postgresql"
	"testing"
)

func CreateUserAndPostForTest(
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

func CreateUserAndGroupForTest(
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