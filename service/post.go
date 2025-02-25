package service

import (
	"fmt"
	"painteer/model"
	"painteer/repository/group"
	post "painteer/repository/post"
)

type PostingService interface {
	CreatePost(uploadPost model.UploadPost) (*model.PostId, error)
	DeletePost(deletePost model.DeletePost) (*model.PostId, error)
	GetPostByID(postId model.PostId) (*model.Post, error)
	GetPostsByPrefectureIDAndGroupIDs(prefectureIDAndGroupIDs model.PrefectureIDAndGroupIDs) ([]model.Post, error)
}

type PostingsServiceImpl struct {
	postRepo  post.PostingsRepository
	groupRepo group.GroupRepository
}

func NewPostingService(postRepo post.PostingsRepository, groupRepo group.GroupRepository) *PostingsServiceImpl {
	return &PostingsServiceImpl{
		postRepo:  postRepo,
		groupRepo: groupRepo,
	}
}

func (s *PostingsServiceImpl) CreatePost(uploadPost model.UploadPost, groupIds []model.GroupId) (*model.PostId, error) {
	createdPost, err := s.postRepo.CreatePost(uploadPost)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	for _, groupId := range groupIds {
		publicSetting := model.PublicSetting{
			PostId:        *createdPost,
			PublicGroupId: groupId,
		}

		if _, err := s.groupRepo.CreatePostPublicSetting(publicSetting); err != nil {
			return nil, fmt.Errorf("failed to set public setting for group %d: %w", groupId, err)
		}
	}

	return createdPost, nil
}

func (s *PostingsServiceImpl) DeletePost(deletePost model.DeletePost) (*model.PostId, error) {
	return s.postRepo.DeletePost(deletePost)
}

func (s *PostingsServiceImpl) GetPostByID(postId model.PostId) (*model.Post, error) {
	return s.postRepo.FindPostByID(postId)
}

func (s *PostingsServiceImpl) GetPostsByPrefectureIDAndGroupIDs(prefectureIDAndGroupIDs model.PrefectureIDAndGroupIDs) ([]model.Post, error) {
	var posts []model.Post

	for _, groupId := range prefectureIDAndGroupIDs.GroupIds {
		prefectureIDAndGroupID := model.PrefectureIDAndGroupID{
			PrefectureId: prefectureIDAndGroupIDs.PrefectureId,
			GroupId:      groupId,
		}

		groupPosts, err := s.postRepo.FindPostsByPrefectureIDAndGroupID(prefectureIDAndGroupID)
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
