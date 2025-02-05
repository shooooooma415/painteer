package service

import (
	"painteer/model"
	"painteer/repository/auth"
)

type UsersServiceImpl struct {
	repo auth.UsersRepository
}

func NewUsersService(repo auth.UsersRepository) *UsersServiceImpl {
	return &UsersServiceImpl{repo: repo}
}

func (s *UsersServiceImpl) RegisterUser(user model.CreateUser) (*model.User, error) {
	return s.repo.CreateUser(user)
}

func (s *UsersServiceImpl) AuthenticateUser(authId model.AuthId) (*model.User, error) {
	return s.repo.SignInUser(authId)
}

func (s *UsersServiceImpl) GetUserByID(userId model.UserId) (*model.User, error) {
	return s.repo.FindUserByUserID(userId)
}
