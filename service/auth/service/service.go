package service

import (
	"painteer/model"
	"painteer/repository/auth"
)

type AuthServiceImpl struct {
	repo auth.UsersRepository
}

func NewAuthService(repo auth.UsersRepository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) RegisterUser(user model.CreateUser) (*model.User, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthServiceImpl) AuthenticateUser(authId model.AuthId) (*model.User, error) {
	return s.repo.SignInUser(authId)
}

func (s *AuthServiceImpl) GetUserByID(userId model.UserId) (*model.User, error) {
	return s.repo.FindUserByUserID(userId)
}
