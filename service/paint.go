package service

import (
	"fmt"
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

func (s *PaintServiceImpl) CountPostIDsByPrefecture(groupIds []model.GroupId) ([]model.CountsByPrefecture, error) {
	prefecturePostMap := make(map[model.PrefectureId]map[model.PostId]struct{})
	for i := 1; i <= 47; i++ {
		prefecturePostMap[model.PrefectureId(i)] = make(map[model.PostId]struct{})
	}

	allPostsByPrefecture := make([]model.PostsByPrefecture, 0)

	for _, groupId := range groupIds {
		postsByPrefecture, err := s.repo.FindPostIDsByPrefecture(groupId)
		if err != nil {
			return nil, fmt.Errorf("failed to find post IDs by prefecture for group %v: %w", groupId, err)
		}
		allPostsByPrefecture = append(allPostsByPrefecture, postsByPrefecture...)
	}

	for _, entry := range allPostsByPrefecture {
		postSet := prefecturePostMap[entry.PrefectureId]
		for _, postId := range entry.PostIds {
			postSet[postId] = struct{}{}
		}
	}

	counts := make([]model.CountsByPrefecture, 0, 47)
	for prefectureId, posts := range prefecturePostMap {
		counts = append(counts, model.CountsByPrefecture{
			Prefecture: model.GetPrefectureName(prefectureId),
			PostCount:    len(posts), 
		})
	}

	return counts, nil
}
