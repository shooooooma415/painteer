package postgresql

import (
	"database/sql"
	"fmt"
	"painteer/model"
	"github.com/lib/pq"
)

type PaintRepositoryImpl struct {
	db *sql.DB
}

func NewPaintRepository(db *sql.DB) *PaintRepositoryImpl {
	return &PaintRepositoryImpl{db: db}
}

func (q *PaintRepositoryImpl) FindPostIDsByPrefecture(groupId model.GroupId) ([]model.PostsByPrefecture, error) {
	query := `
		SELECT p.prefecture_id, ARRAY_AGG(p.id) AS post_ids
		FROM posts p
		INNER JOIN public_setting ps ON p.id = ps.post_id
		WHERE ps.group_id = $1
		GROUP BY p.prefecture_id
		ORDER BY p.prefecture_id
	`

	rows, err := q.db.Query(query, groupId)
	if err != nil {
		return nil, fmt.Errorf("failed to find post IDs by prefecture for group_id %v: %w", groupId, err)
	}
	defer rows.Close()

	var result []model.PostsByPrefecture

	for rows.Next() {
		var postsByPrefecture model.PostsByPrefecture
		if err := rows.Scan(&postsByPrefecture.PrefectureId, pq.Array(&postsByPrefecture.PostIds)); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		result = append(result, postsByPrefecture)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return result, nil
}

