package postgresql

import (
	"database/sql"
	"fmt"
	"painteer/model"
)

type GroupRepositoryImpl struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{db: db}
}

func (q *GroupRepositoryImpl) CreateGroup(createGroup model.CreateGroup) (*model.Group, error) {
	query := `
		INSERT INTO groups (name, password, author_id)
		VALUES ($1, $2, $3)
		RETURNING *
	`
	var createdGroup model.Group

	err := q.db.QueryRow(
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

func (q *GroupRepositoryImpl) JoinGroup(joinGroup model.JoinGroup) (*model.JoinGroup, error) {
	query := `
		INSERT INTO user_group (group_id, user_id)
		VALUES ($1, $2)
		RETURNING group_id, user_id
	`

	var joinedGroup model.JoinGroup

	err := q.db.QueryRow(
		query,
		joinGroup.GroupId,
		joinGroup.UserId,
	).Scan(
		&joinedGroup.GroupId,
		&joinedGroup.UserId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create group:%w", err)
	}
	return &joinedGroup, nil
}

