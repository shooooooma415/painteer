package group

import "painteer/model"

type GroupRepository interface {
	CreateGroup(createGroup model.Group) (*model.Group, error)
	VerifyPassword(VerifyPassword model.VerifyPassword)(*model.GroupId,error)
	IsUserExist(joinGroup model.JoinGroup)(bool,error)
	JoinGroup(joinGroup model.JoinGroup) (*model.Group, error)
	FetchGroup(userId model.UserId) (*model.FetchedGroup, error)
	InsertPublicSetting(publicSetting model.PublicSetting) (*model.PublicSetting, error)
}
