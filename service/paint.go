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
	for _, prefectureId := range model.AllPrefectureIds {
		prefecturePostMap[prefectureId] = make(map[model.PostId]struct{})
	}

	for _, groupId := range groupIds {
		postsByPrefecture, err := s.repo.FindPostIDsByPrefecture(groupId)
		if err != nil {
			return nil, fmt.Errorf("failed to find post IDs by prefecture for group %v: %w", groupId, err)
		}

		for _, entry := range postsByPrefecture {
			for _, postId := range entry.PostIds {
				prefecturePostMap[entry.PrefectureId][postId] = struct{}{}
			}
		}
	}

	counts := make([]model.CountsByPrefecture, 0, len(model.AllPrefectureIds))
	for _, prefectureId := range model.AllPrefectureIds {
		counts = append(counts, model.CountsByPrefecture{
			Prefecture: model.GetPrefectureName(prefectureId),
			PostCount:  len(prefecturePostMap[prefectureId]),
		})
	}

	return counts, nil
}


func (s *PaintServiceImpl) CountPostIDsByRegion(groupIds []model.GroupId) ([]model.CountsByRegion, error) {
	prefectureCounts, err := s.CountPostIDsByPrefecture(groupIds)
	if err != nil {
		return nil, fmt.Errorf("failed to count posts by prefecture: %w", err)
	}

	regionMap := map[string][]model.PrefectureId{
		"Hokkaido":       {model.Hokkaido},
		"Tohoku":         {model.Aomori, model.Iwate, model.Miyagi, model.Akita, model.Yamagata, model.Fukushima},
		"Kanto":          {model.Ibaraki, model.Tochigi, model.Gunma, model.Saitama, model.Chiba, model.Tokyo, model.Kanagawa},
		"Chubu":          {model.Niigata, model.Toyama, model.Ishikawa, model.Fukui, model.Yamanashi, model.Nagano, model.Gifu, model.Shizuoka, model.Aichi},
		"Kinki":          {model.Mie, model.Shiga, model.Kyoto, model.Osaka, model.Hyogo, model.Nara, model.Wakayama},
		"Chugoku":        {model.Tottori, model.Shimane, model.Okayama, model.Hiroshima, model.Yamaguchi},
		"Shikoku":        {model.Tokushima, model.Kagawa, model.Ehime, model.Kochi},
		"Kyushu/Okinawa": {model.Fukuoka, model.Saga, model.Nagasaki, model.Kumamoto, model.Oita, model.Miyazaki, model.Kagoshima, model.Okinawa},
	}

	regionCounts := make(map[string]int)

	for _, count := range prefectureCounts {
		for region, prefectures := range regionMap {
			for _, prefectureId := range prefectures {
				if count.Prefecture == model.GetPrefectureName(prefectureId) {
					regionCounts[region] += count.PostCount
					break
				}
			}
		}
	}

	countsByRegion := make([]model.CountsByRegion, 0, len(regionCounts))
	for region, count := range regionCounts {
		countsByRegion = append(countsByRegion, model.CountsByRegion{
			Region:    region,
			PostCount: count,
		})
	}

	return countsByRegion, nil
}
