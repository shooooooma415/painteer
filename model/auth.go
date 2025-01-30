package model

type UserName string

type UserId int

type AuthId string

type CreateUser struct {
	UserName UserName
	Icon     string
	AuthId   AuthId
}

type User struct {
	UserName UserName
	Icon     string
	AuthId   AuthId
	UserId   UserId
}
