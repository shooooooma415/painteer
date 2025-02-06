package posting

import "painteer/model"

type PostingService interface {
	CreatePost(uploadPost model.UploadPost) (*model.PostId, error)
	DeletePost(postId model.PostId) (*model.PostId, error)
	GetPostByID(postId model.PostId) (*model.Post, error)
	GetPostsByPrefectureIDAndGroupIDs(prefectureIDAndGroupIDs model.PrefectureIDAndGroupIDs)(*model.Posts,error)
}
