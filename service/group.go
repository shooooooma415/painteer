package service

import (
	"painteer/model"
	"painteer/repository/group"
)

type GroupService interface {
	RegisterGroup(Group model.CreateGroup) (*model.Group, error)
	JoinGroup(CreateUserGroup model.CreateUserGroup) (*model.CreateUserGroup, error)
	GetUserGroupSummaryByUserID(userId model.UserId) ([]model.GroupSummary, error)
	GetGroupMembersByGroupID(groupId model.GroupId) (*model.GroupMembers, error)
	GetGroupSummaryByGroupID(groupId model.GroupId) (*model.GroupSummary, error)
	RegisterPublicSetting(publicSetting model.PublicSetting) (*model.PublicSetting, error)
}

type GroupServiceImpl struct {
	repo group.GroupRepository
}

func NewAuthService(repo group.GroupRepository) *GroupServiceImpl {
	return &GroupServiceImpl{repo: repo}
}

func (s *GroupServiceImpl) RegisterGroup(Group model.CreateGroup) (*model.Group, error) {
	return s.repo.CreateGroup(Group)
}

func (s *GroupServiceImpl) JoinGroup(CreateUserGroup model.CreateUserGroup) (*model.CreateUserGroup, error) {

}

func (s *GroupServiceImpl) GetUserGroupSummaryByUserID(userId model.UserId) ([]model.GroupSummary, error) {

}

func (s *GroupServiceImpl) GetGroupMembersByGroupID(groupId model.GroupId) (*model.GroupMembers, error) {
	return s.repo.FindGroupMembersByGroupID(groupId)
}

func (s *GroupServiceImpl) RegisterPublicSetting(publicSetting model.PublicSetting) (*model.PublicSetting, error) {
	return s.repo.CreatePublicSetting(publicSetting)
}
