package postgresql

import (
	"database/sql"
	"fmt"
	"painteer/model"
)

type GroupRepositoryImpl struct {
	DB *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{DB: db}
}

func (q *GroupRepositoryImpl) CreateGroup(createGroup model.CreateGroup) (*model.Group, error) {
	query := `
		INSERT INTO groups (name, password, author_id)
		VALUES ($1, $2, $3)
		RETURNING *
	`
	var createdGroup model.Group

	err := q.DB.QueryRow(
		query,
		createGroup.GroupName,
		createGroup.Password,
		createGroup.AuthorId,
	).Scan(
		&createdGroup.GroupId,
		&createdGroup.GroupName,
		&createdGroup.Icon,
		&createdGroup.Password,
		&createdGroup.AuthorId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create group:%w", err)
	}
	return &createdGroup, nil
}
