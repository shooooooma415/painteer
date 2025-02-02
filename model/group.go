package model

type GroupId int

type GroupName string

type Password string

type Group struct {
	GroupName GroupName
	Icon      string
	AuthorId  UserId
	GroupId   GroupId
}

type CreateGroup struct {
	GroupName GroupName
	Icon      string
	Password  Password
	AuthorId  UserId
}

type JoinGroup struct {
	UserId    UserId
	GroupName GroupName
	Password  Password
}

type FetchedGroup struct {
	Groups []Group
}

type PublicSetting struct {
	PostId       PostId
	publicGroups []GroupId
}
