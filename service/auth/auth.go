package auth

import "painteer/model"

type AuthService interface {
	RegisterUser(user model.CreateUser) (*model.User, error)
	AuthenticateUser(authId model.AuthId) (*model.User, error)
	GetUserByID(userId model.UserId) (*model.User, error)
}
