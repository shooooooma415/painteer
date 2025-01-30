package post

import "painteer/model"

type PostingsRepository interface {
	UploadPost(uploadPost model.UploadPost) (*model.Post, error)
	DeletePost(postId model.PostId) (*model.Post, error)
	SelectPost(selectPost model.SelectPost)(model.Post,error)
}