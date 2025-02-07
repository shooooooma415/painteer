package service

import "painteer/model"

type GroupService interface {
	RegisterGroup(createGroup model.Group) (*model.Group, error)
	JoinGroup(CreateUserGroup model.CreateUserGroup) (*model.CreateUserGroup, error)
	GetUserGroupSummaryByUserID(userId model.UserId) ([]model.GroupSummary, error)
	GetGroupMembersByGroupID(groupId model.GroupId) (*model.GroupMembers, error)
	GetGroupSummaryByGroupID(groupId model.GroupId) (*model.GroupSummary, error)
	RegisterPublicSetting(publicSetting model.PublicSetting) (*model.PublicSetting, error)
}
