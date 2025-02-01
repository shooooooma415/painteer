package postgresql

import (
	"database/sql"
	"fmt"
	"painteer/model"
)

type PostRepositoryImpl struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{DB: db}
}

func (q *PostRepositoryImpl) CreatePost(uploadPost model.UploadPost) (*model.PostId, error) {
	query := `
		INSERT INTO posts (
				image, comment, prefecture_id, user_id, date, longitude, latitude
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var uploadedPostId model.PostId
	err := q.DB.QueryRow(
		query,
		uploadPost.Image,
		uploadPost.Comment,
		uploadPost.PrefectureId,
		uploadPost.UserId,
		uploadPost.Date,
		uploadPost.Longitude,
		uploadPost.Latitude,
	).Scan(
		&uploadedPostId,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to upload post: %w", err)
	}

	return &uploadedPostId, nil
}

func (q *PostRepositoryImpl) DeletePost(deletePost model.DeletePost) (*model.PostId, error) {
	query := `
		DELETE FROM posts
		WHERE id = $1 AND user_id = $2
		RETURNING id
		`

	var deletedPostId model.PostId
	err := q.DB.QueryRow(query, deletePost.PostId, deletePost.UserId).Scan(
		&deletedPostId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete post: %w", err)
	}

	return &deletedPostId, nil
}
