package service

import (
	"painteer/model"
	"painteer/repository/posting"
)

type PostingsServiceImpl struct {
	repo post.PostingsRepository
}

func NewPostingService(repo post.PostingsRepository) *PostingsServiceImpl {
	return &PostingsServiceImpl{repo: repo}
}

func (s *PostingsServiceImpl) CreatePost(uploadPost model.UploadPost) (*model.PostId, error){
	return s.repo.CreatePost(uploadPost)
}

func (s *PostingsServiceImpl)DeletePost(postId model.PostId) (*model.PostId, error){
	return s.repo.DeletePost(postId)
}

func (s *PostingsServiceImpl)GetPostByID(postId model.PostId) (*model.Post, error){
	return s.repo.FindPostByID(postId)
}
