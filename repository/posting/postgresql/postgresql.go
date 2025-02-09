package postgresql

import (
	"database/sql"
	"fmt"
	"painteer/model"
)

type PostRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{db: db}
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
	err := q.db.QueryRow(
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
	err := q.db.QueryRow(query, deletePost.PostId, deletePost.UserId).Scan(
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
	err := q.db.QueryRow(query, postId).Scan(
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

func (q *PostRepositoryImpl) FindPostsByPrefectureIDAndGroupID(prefectureIDAndGroupID model.PrefectureIDAndGroupID) ([]model.Post, error) {
	query := `
		SELECT p.id, p.image, p.comment, p.prefecture_id, p.user_id, p.date, p.longitude, p.latitude
		FROM posts p
		INNER JOIN public_setting ps ON p.id = ps.post_id
		WHERE p.prefecture_id = $1
		AND ps.group_id = $2
	`

	rows, err := q.db.Query(
		query,
		prefectureIDAndGroupID.PrefectureId,
		prefectureIDAndGroupID.GroupId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find posts by prefecture_id and group_id: %w", err)
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(
			&post.PostId,
			&post.Image,
			&post.Comment,
			&post.PrefectureId,
			&post.UserId,
			&post.Date,
			&post.Longitude,
			&post.Latitude,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return posts, nil
}

