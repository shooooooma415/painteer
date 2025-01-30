package auth

import "painteer/model"

type PostingsRepository interface {
	CreatePosting(user model.CreateUser) (*model.User, error)
	SignInUser(authId model.AuthId) (*model.User, error)
}