package auth

import "painteer/model"

type UsersRepository interface {
	CreateUser(user model.CreateUser) (*model.UserId, error)
	SignInUser(authId model.AuthId) (*model.UserId, error)
}