package model

type GroupId int

type GroupName string

type Password string

type Group struct {
	GroupName GroupName
	Icon      string
	Password  Password
	UserId    UserId
	GroupId   GroupId
}

type CreateGroup struct {
	GroupName GroupName
	Icon      string
	Password  Password
	UserId    UserId
}

type JoinGroup struct {
	UserId    UserId
	GroupName GroupName
	Password  Password
}

type FetchedGroup struct {
	Groups []Group
}
