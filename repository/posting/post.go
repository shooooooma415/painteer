package post

import "painteer/model"

type PostingsRepository interface {
	CreatePost(uploadPost model.UploadPost) (*model.PostId, error)
	DeletePost(postId model.PostId) (*model.PostId, error)
	FindPostByID(postId model.PostId) (*model.Post, error)
}
