package group

import "painteer/model"

type GroupRepository interface {
	CreateGroup(createGroup model.Group) (*model.Group, error)
	FindPasswordAndGroupName() (*model.PasswordAndNames, error)
	InsertGroup(insertGroup model.InsertGroup) (*model.Group, error)
	FindUserGroupsByUserID(userId model.UserId) (*model.UserGroups, error)
	FindGroupByGroupID(groupId model.GroupId) (*model.Group, error)
	FindGroupMembersByGroupID(groupId model.GroupId) (*model.GroupMembers, error)
	CreatePublicSetting(publicSetting model.PublicSetting) (*model.PublicSetting, error)
}
