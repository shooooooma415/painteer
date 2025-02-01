package postgresql

import (
	"database/sql"
	"fmt"
	"painteer/model"

	"github.com/lib/pq"
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
		WHERE id = $1
		RETURNING id
		`

	var deletedPost model.Post
	err := q.DB.QueryRow(query, postId).Scan(
		&deletedPost.PostId,
		&deletedPost.Image,
		&deletedPost.Comment,
		&deletedPost.PrefectureId,
		&deletedPost.UserId,
		&deletedPost.Date,
		&deletedPost.Longitude,
		&deletedPost.Latitude,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete post: %w", err)
	}

	return &deletedPost, nil
}

func (q *PostRepositoryImpl) SelectPost(selectPost model.SelectPost) (*model.Post, error) {
	query := `
		SELECT 
			p.id, p.image, p.comment, p.prefecture_id, 
			p.user_id, p.date, p.longitude, p.latitude
		FROM posts p
		INNER JOIN public_setting ps
		ON p.id = ps.post_id
		WHERE p.prefecture_id = $1
		AND ps.group_id = ANY($2::int[])
	`

	var selectedPost model.Post
	err := q.DB.QueryRow(
		query,
		selectPost.PrefectureId,
		pq.Array(selectPost.Groups), // `[]int` を PostgreSQL の `int[]` 型として渡す
	).Scan(
		&selectedPost.PostId,
		&selectedPost.Image,
		&selectedPost.Comment,
		&selectedPost.PrefectureId,
		&selectedPost.UserId,
		&selectedPost.Date,
		&selectedPost.Longitude,
		&selectedPost.Latitude,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to select post: %w", err)
	}

	return &selectedPost, nil
}
