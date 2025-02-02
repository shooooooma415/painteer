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

type JoinGroup struct {
	UserId    UserId
	GroupName GroupName
	Password  Password
}
