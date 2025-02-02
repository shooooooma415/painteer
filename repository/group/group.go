package group

import "painteer/model"

type GroupRepository interface {
	createGroup(createGroup model.Group) (*model.Group, error)
	JoinGroup(joinGroup model.JoinGroup) (*model.Group, error)
	
}
