package postgresql

import (
	"database/sql"
	"fmt"
	"painteer/model"
)

type PostRepositoryImpl struct {
	dB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{DB: db}
}

func (q *PostRepositoryImpl) CreatePost(uploadPost model.UploadPost) (*model.Post, error) {
	query := `
		INSERT INTO posts (
				image, comment, prefecture_id, user_id, date, longitude, latitude
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, image, comment, prefecture_id, user_id, date, longitude, latitude
	`

	var uploadedPost model.Post
	err := q.dB.QueryRow(
		query,
		uploadPost.Image,
		uploadPost.Comment,
		uploadPost.PrefectureId,
		uploadPost.UserId,
		uploadPost.Date,
		uploadPost.Longitude,
		uploadPost.Latitude,
	).Scan(
		&uploadedPost.PostId,
		&uploadedPost.Image,
		&uploadedPost.Comment,
		&uploadedPost.PrefectureId,
		&uploadedPost.UserId,
		&uploadedPost.Date,
		&uploadedPost.Longitude,
		&uploadedPost.Latitude,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to upload post: %w", err)
	}

	return &uploadedPost, nil
}

func (q *PostRepositoryImpl) DeletePost(deletePost model.DeletePost) (*model.PostId, error) {
	query := `
		DELETE FROM posts
		WHERE id = $1 AND user_id = $2
		RETURNING id
		`

	var deletedPostId model.PostId
	err := q.dB.QueryRow(query, deletePost.PostId, deletePost.UserId).Scan(
		&deletedPostId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete post: %w", err)
	}

	return &deletedPostId, nil
}

func (q *PostRepositoryImpl) FindPostByID(postId model.PostId) (*model.Post, error) {
	query := `
		SELECT *
		FROM posts
		WHERE id = $1
	`

	var fetchedPost model.Post
	err := q.dB.QueryRow(query, postId).Scan(
		&fetchedPost.PostId,
		&fetchedPost.Image,
		&fetchedPost.Comment,
		&fetchedPost.PrefectureId,
		&fetchedPost.UserId,
		&fetchedPost.Date,
		&fetchedPost.Longitude,
		&fetchedPost.Latitude,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete post: %w", err)
	}

	return &fetchedPost, nil
}
