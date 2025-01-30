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

func (q *PostRepositoryImpl) UploadPost(uploadPost model.UploadPost) (*model.Post, error) {
	query := `
		WITH uploaded_post AS (
			INSERT INTO posts (
				image, comment, prefecture_id, user_id, date, longitude, latitude
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id
		)
		INSERT INTO public_setting (post_id, group_id)
		SELECT uploaded_post.id, unnest($8::int[])
		FROM uploaded_post
		RETURNING *;
	`

	var uploadedPost model.Post
	err := q.DB.QueryRow(
		query,
		uploadPost.Image,
		uploadPost.Comment,
		uploadPost.PrefectureId,
		uploadPost.UserID,
		uploadPost.Date,
		uploadPost.Longitude,
		uploadPost.Latitude,
		pq.Array(uploadPost.Groups), // `[]GroupId` を PostgreSQL の int[]型として渡せるらしい
	).Scan(
		&uploadedPost.PostId,
		&uploadedPost.Image,
		&uploadedPost.Comment,
		&uploadedPost.PrefectureId,
		&uploadedPost.UserID,
		&uploadedPost.Date,
		&uploadedPost.Longitude,
		&uploadedPost.Latitude,
		&uploadedPost.Groups,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to upload post: %w", err)
	}

	return &uploadedPost, nil
}

func (q *PostRepositoryImpl) DeletePost(postId model.PostId) (*model.Post, error) {
	query := `
		WITH deleted_posts AS (
			DELETE FROM posts
			WHERE id = $1
			RETURNING id, image, comment, prefecture_id, user_id, date, longitude, latitude
		), deleted_settings AS (
			DELETE FROM public_setting
			WHERE post_id IN (SELECT id FROM deleted_posts)
		)
		SELECT id, image, comment, prefecture_id, user_id, date, longitude, latitude FROM deleted_posts;
	`

	var deletedPost model.Post
	err := q.DB.QueryRow(query, postId).Scan(
		&deletedPost.PostId,
		&deletedPost.Image,
		&deletedPost.Comment,
		&deletedPost.PrefectureId,
		&deletedPost.UserID,
		&deletedPost.Date,
		&deletedPost.Longitude,
		&deletedPost.Latitude,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete post: %w", err)
	}

	return &deletedPost, nil
}
