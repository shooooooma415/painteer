package post

import "painteer/model"

type PostingsRepository interface {
	CreatePost(uploadPost model.UploadPost) (*model.PostId, error)
	DeletePost(deletePost model.DeletePost) (*model.PostId, error)
	FindPostByID(postId model.PostId) (*model.Post, error)
	FindPostsByPrefectureIDAndGroupID(prefectureIDAndGroupID model.PrefectureIDAndGroupID)([]model.Post,error)
}
