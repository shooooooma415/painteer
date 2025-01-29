package model

type UserName string

type UserId string

type AuthId string

type CreateUser struct {
	UserName UserName
	Icon     string
	AuthId   AuthId
}
