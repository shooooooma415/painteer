package auth

import "painteer/model"

type UsersRepository interface {
	CreateUser(user model.CreateUser) (*model.User, error)
	SignInUser(authId model.AuthId) (*model.User, error)
}