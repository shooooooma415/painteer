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

type VerifyPassword struct {
	Password  Password
	GroupName GroupName
}

type JoinGroup struct {
	UserId  UserId
	GroupId GroupId
}

type FetchedGroups struct {
	Groups []Group
}

type PublicSetting struct {
	PostId        PostId
	PublicGroupId GroupId
}

type GroupMember struct {
	GroupId GroupId
	Members []UserName
}

type PostId string //マージしたら消す