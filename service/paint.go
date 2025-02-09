package service

import (
	"painteer/model"
	"painteer/repository/paint"
)

type PaintService interface {
	RegisterUser(user model.CreateUser) (*model.User, error)
	AuthenticateUser(authId model.AuthId) (*model.User, error)
	GetUserByID(userId model.UserId) (*model.User, error)
}

type PaintServiceImpl struct {
	repo paint.PaintRepository
}

func NewPaintService(repo paint.PaintRepository) *PaintServiceImpl {
	return &PaintServiceImpl{repo: repo}
}
