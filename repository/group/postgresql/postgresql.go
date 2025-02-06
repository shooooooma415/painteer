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
		INSERT INTO groups (name, password, author_id, icon)
		VALUES ($1, $2, $3, $4)
		RETURNING *
	`
	var createdGroup model.Group

	err := q.db.QueryRow(
		query,
		createGroup.GroupName,
		createGroup.Password,
		createGroup.AuthorId,
		createGroup.Icon,
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

func (q *GroupRepositoryImpl) InsertGroup(InsertGroup model.InsertGroup) (*model.InsertGroup, error) {
	query := `
		INSERT INTO user_group (group_id, user_id)
		VALUES ($1, $2)
		RETURNING group_id, user_id
	`

	var joinedGroup model.InsertGroup

	err := q.db.QueryRow(
		query,
		InsertGroup.GroupId,
		InsertGroup.UserId,
	).Scan(
		&joinedGroup.GroupId,
		&joinedGroup.UserId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to join group:%w", err)
	}
	return &joinedGroup, nil
}

func (q *GroupRepositoryImpl) FindGroupByGroupID(groupId model.GroupId) (*model.Group, error) {
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

func (q *GroupRepositoryImpl) FindGroupMembersByGroupID(groupId model.GroupId) (*model.GroupMembers, error) {
	query := `
		SELECT ug.group_id, u.id
		FROM user_group ug
		LEFT JOIN users u
		ON ug.user_id = u.id
		WHERE ug.group_id = $1
	`

	rows, err := q.db.Query(query, groupId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch group members: %w", err)
	}
	defer rows.Close()

	var members []model.UserId
	var fetchedGroupId model.GroupId

	firstRow := true
	for rows.Next() {
		var userId model.UserId
		var currentGroupId model.GroupId

		if err := rows.Scan(&currentGroupId, &userId); err != nil {
			return nil, fmt.Errorf("failed to scan user_name: %w", err)
		}

		if firstRow {
			fetchedGroupId = currentGroupId
			firstRow = false
		}

		members = append(members, userId)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if firstRow {
		return nil, fmt.Errorf("no members found for group_id: %d", groupId)
	}

	returnValue := model.GroupMembers{
		GroupId: fetchedGroupId,
		Members: members,
	}

	return &returnValue, nil
}

func (q *GroupRepositoryImpl) CreatePostPublicSetting(ps model.PublicSetting) (*model.PublicSetting, error) {
	query := `
		INSERT INTO public_setting (post_id, group_id)
		VALUE $1, $2
		RETURNING post_id, group_id
	`
	var setPost model.PublicSetting
	err := q.db.QueryRow(query, ps.PostId, ps.PostId).Scan(&setPost.PostId, &setPost.PublicGroupId)

	if err != nil {
		return nil, fmt.Errorf("failed to create public setting: %w", err)
	}
	return &setPost, nil
}

func (q *GroupRepositoryImpl) FindUserGroupsByUserID(userId model.UserId) (*model.UserGroups, error) {
	query := `
		SELECT ug.user_id, ug.group_id, g.name, g.icon, g.password, g.author_id
		FROM user_group ug
		LEFT JOIN groups g
		ON ug.group_id = g.id
		WHERE ug.user_id = $1
	`

	rows, err := q.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch groups for user: %w", err)
	}
	defer rows.Close()

	var fetchedGroups model.UserGroups
	fetchedGroups.UserId = userId
	var groups []model.Group

	for rows.Next() {
		var group model.Group
		err := rows.Scan(
			&fetchedGroups.UserId,
			&group.GroupId,
			&group.GroupName,
			&group.Icon,
			&group.Password,
			&group.AuthorId,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan group data: %w", err)
		}
		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	fetchedGroups.Groups = groups
	return &fetchedGroups, nil
}
