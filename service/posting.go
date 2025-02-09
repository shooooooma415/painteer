package service

import (
	"fmt"
	"painteer/model"
	post "painteer/repository/posting"
)

type PostingService interface {
	CreatePost(uploadPost model.UploadPost) (*model.PostId, error)
	DeletePost(postId model.PostId) (*model.PostId, error)
	GetPostByID(postId model.PostId) (*model.Post, error)
	GetPostsByPrefectureIDAndGroupIDs(prefectureIDAndGroupIDs model.PrefectureIDAndGroupIDs) ([]model.Post, error)
}

type PostingsServiceImpl struct {
	repo post.PostingsRepository
}

func NewPostingService(repo post.PostingsRepository) *PostingsServiceImpl {
	return &PostingsServiceImpl{repo: repo}
}

func (s *PostingsServiceImpl) CreatePost(uploadPost model.UploadPost) (*model.PostId, error) {
	return s.repo.CreatePost(uploadPost)
}

func (s *PostingsServiceImpl) DeletePost(postId model.PostId) (*model.PostId, error) {
	return s.repo.DeletePost(postId)
}

func (s *PostingsServiceImpl) GetPostByID(postId model.PostId) (*model.Post, error) {
	return s.repo.FindPostByID(postId)
}

func (s *PostingsServiceImpl) GetPostsByPrefectureIDAndGroupIDs(prefectureIDAndGroupIDs model.PrefectureIDAndGroupIDs) ([]model.Post, error) {
	var posts []model.Post

	for _, groupId := range prefectureIDAndGroupIDs.GroupIds {
		prefectureIDAndGroupID := model.PrefectureIDAndGroupID{
			PrefectureId: prefectureIDAndGroupIDs.PrefectureId,
			GroupId:      groupId,
		}

		groupPosts, err := s.repo.FindPostsByPrefectureIDAndGroupID(prefectureIDAndGroupID)
		if err != nil {
			fmt.Printf("Error fetching posts for PrefectureId %v and GroupId %v: %v\n", prefectureIDAndGroupIDs.PrefectureId, groupId, err)
			continue
		}

		if len(groupPosts) > 0 {
			posts = append(posts, groupPosts...)
		}
	}

	return posts, nil
}

