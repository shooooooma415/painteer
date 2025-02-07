package service

import (
	"painteer/model"
	"painteer/repository/auth"
)

type AuthService interface {
	RegisterUser(user model.CreateUser) (*model.User, error)
	AuthenticateUser(authId model.AuthId) (*model.User, error)
	GetUserByID(userId model.UserId) (*model.User, error)
}

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
	return s.repo.FindUserByAuthID(authId)
}

func (s *AuthServiceImpl) GetUserByID(userId model.UserId) (*model.User, error) {
	return s.repo.FindUserByUserID(userId)
}
