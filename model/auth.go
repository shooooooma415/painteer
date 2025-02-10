package model

type UserName string

type UserId int

type AuthId string

type CreateUser struct {
	UserName string `json:"name"`
	Icon     string `json:"icon"`
	AuthId   string `json:"auth_id"`
}

type User struct {
	UserName UserName
	Icon     string
	AuthId   AuthId
	UserId   UserId
}

type SignUpResponse struct {
	UserId UserId `json:"user_id"`
}
