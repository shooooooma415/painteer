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

type CreateUserGroup struct {
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

type GroupSummary struct {
	GroupId   GroupId
	GroupName GroupName
	Icon      string
}

type JoinGroup struct {
	UserId   UserId
	GroupId  GroupId
	Password Password
}

type GetGroupMembersResponse struct {
	Member []struct {
		UserId   UserId   `json:"user_id"`
		UserName UserName `json:"user_name"`
		Icon     string   `json:"icon"`
	} `json:"member"`
}

type GetUserGroupResponse struct {
	Groups []GroupSummary `json:"groups"`
}
