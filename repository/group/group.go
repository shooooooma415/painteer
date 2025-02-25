package group

import "painteer/model"

type GroupRepository interface {
	CreateGroup(createGroup model.CreateGroup) (*model.Group, error)
	CreateUserGroup(CreateUserGroup model.CreateUserGroup) (*model.CreateUserGroup, error)
	FindUserGroupsByUserID(userId model.UserId) (*model.UserGroups, error)
	FindGroupByGroupID(groupId model.GroupId) (*model.Group, error)
	FindGroupMembersByGroupID(groupId model.GroupId) (*model.GroupMembers, error)
	CreatePostPublicSetting(publicSetting model.PublicSetting) (*model.PublicSetting, error)
}
