package model

type GroupId int

type GroupName string

type Password string

type Group struct {
	GroupId   GroupId
	GroupName GroupName
	Icon      string
	Password  Password
	AuthorId  UserId
}

type CreateGroup struct {
	GroupName GroupName
	Icon      string
	Password  Password
	AuthorId  UserId
}

type PasswordAndName struct {
	Password  Password
	GroupName GroupName
}

type InsertGroup struct {
	UserId  UserId
	GroupId GroupId
}

type UserGroups struct {
	UserId UserId
	Groups []Group
}

type PublicSetting struct {
	PostId        PostId
	PublicGroupId GroupId
}

type GroupMembers struct {
	GroupId GroupId
	Members []UserId
}

type PasswordAndNames struct {
	Groups []PasswordAndName
}
