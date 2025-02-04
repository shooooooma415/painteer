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
		return nil, fmt.Errorf("failed to join group:%w", err)
	}
	return &joinedGroup, nil
}

func (q *GroupRepositoryImpl) FindGroupIDByPasswordAndName(verifyPassword model.VerifyPassword) (*model.GroupId, error) {
	query := `
			SELECT id FROM groups
			WHERE password = $1
			AND name = $2
	`
	var groupId model.GroupId
	err := q.db.QueryRow(
		query,
		verifyPassword.Password,
		verifyPassword.GroupName,
	).Scan(
		&groupId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("group not found: password=%s, name=%s", verifyPassword.Password, verifyPassword.GroupName)
		}
		return nil, fmt.Errorf("failed to join group: %w", err)
	}

	return &groupId, nil
}

func (q *GroupRepositoryImpl) IsUserExist(joinGroup model.JoinGroup) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM user_group
			WHERE user_id = $1
			AND group_id = $2
		)
	`

	var isExist bool
	err := q.db.QueryRow(query, joinGroup.UserId, joinGroup.GroupId).Scan(&isExist)
	if err != nil {
		return false, fmt.Errorf("failed to check user existence in group: %w", err)
	}

	return isExist, nil
}

func (q *GroupRepositoryImpl) FindGroupByID(groupId model.GroupId) (*model.Group, error) {
	query := `
		SELECT id, name, icon, password, author_id
		FROM groups
		WHERE id = $1
	`

	var foundGroup model.Group
	err := q.db.QueryRow(query, groupId).Scan(
		&foundGroup.GroupId,
		&foundGroup.GroupName,
		&foundGroup.Icon,
		&foundGroup.Password,
		&foundGroup.AuthorId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find group: %w", err)
	}
	return &foundGroup, nil
}
