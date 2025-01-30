package post

import "painteer/model"

type PostingsRepository interface {
	UploadPost(uploadPost model.UploadPost) (*model.Post, error)
	DeletePost(postId model.PostId) (*model.PostId, error)
	SelectPost(selectPost model.SelectPost)(model.Post,error)
}