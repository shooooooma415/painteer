package service

import (
	"fmt"
	"painteer/model"
	"painteer/repository/group"
)

type GroupService interface {
	RegisterGroup(Group model.CreateGroup) (*model.Group, error)
	JoinGroup(JoinGroup model.JoinGroup) (*model.GroupId, error)
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

func (s *GroupServiceImpl) JoinGroup(joinGroup model.JoinGroup) (*model.GroupId, error) {
	group, err := s.repo.FindGroupByGroupID(joinGroup.GroupId)
	if err != nil {
		return nil, fmt.Errorf("failed to find group: %w", err)
	}

	if group == nil {
		return nil, fmt.Errorf("group not found")
	}

	if group.Password != joinGroup.Password {
		return nil, fmt.Errorf("incorrect password")
	}

	userGroups, err := s.repo.FindUserGroupsByUserID(joinGroup.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user groups: %w", err)
	}

	for _, g := range userGroups.Groups {
		if g.GroupId == joinGroup.GroupId {
			return nil, fmt.Errorf("user already in group")
		}
	}

	userGroup := model.CreateUserGroup{
		UserId:  joinGroup.UserId,
		GroupId: joinGroup.GroupId,
	}

	_, err = s.repo.CreateUserGroup(userGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to join group: %w", err)
	}

	return &joinGroup.GroupId, nil
}

func (s *GroupServiceImpl) GetUserGroupSummaryByUserID(userId model.UserId) ([]model.GroupSummary, error) {

}

func (s *GroupServiceImpl) GetGroupMembersByGroupID(groupId model.GroupId) (*model.GroupMembers, error) {
	return s.repo.FindGroupMembersByGroupID(groupId)
}

func (s *GroupServiceImpl) RegisterPublicSetting(publicSetting model.PublicSetting) (*model.PublicSetting, error) {
	return s.repo.CreatePublicSetting(publicSetting)
}
