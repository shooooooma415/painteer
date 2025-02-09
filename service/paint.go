package service

import (
	"painteer/model"
	"painteer/repository/paint"
)

type PaintService interface {
	CountPostIDsByPrefecture(groupIds []model.GroupId) ([]model.CountsByPrefecture, error)
	CountPostIDsByRegion(groupIds []model.GroupId) ([]model.CountsByRegion, error)

}

type PaintServiceImpl struct {
	repo paint.PaintRepository
}

func NewPaintService(repo paint.PaintRepository) *PaintServiceImpl {
	return &PaintServiceImpl{repo: repo}
}


