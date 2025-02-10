package auth

import "painteer/model"

type UsersRepository interface {
	CreateUser(user model.CreateUser) (*model.User, error)
	FindUserByAuthID(authId model.AuthId) (*model.User, error)
	FindUserByUserID(userId model.UserId)(*model.User, error)
}