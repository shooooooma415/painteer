package postgresql

import (
	"database/sql"
	"fmt"
	"painteer/model"
)

type PaintRepositoryImpl struct {
	db *sql.DB
}

func NewPaintRepository(db *sql.DB) *PaintRepositoryImpl {
	return &PaintRepositoryImpl{db: db}
}

func (q *PaintRepositoryImpl) CountPostsByPrefecture(groupId model.GroupId) (*model.Count, error) {
	query := `
		SELECT p.prefecture_id, COUNT(p.id) AS post_count
		FROM posts p
		INNER JOIN public_setting ps
		ON p.id = ps.post_id
		WHERE ps.group_id = $1
		GROUP BY p.prefecture_id
		ORDER BY p.prefecture_id
`

	rows, err := q.db.Query(query, groupId)
	if err != nil {
		return nil, fmt.Errorf("failed to count posts by prefecture for group_id %v: %w", groupId, err)
	}
	defer rows.Close()

	counts := model.Count{}

	for rows.Next() {
		var countByPrefecture model.CountByPrefectureID
		if err := rows.Scan(&countByPrefecture.PrefectureId, &countByPrefecture.PostCount); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		counts.Data = append(counts.Data, countByPrefecture)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return &counts, nil
}
