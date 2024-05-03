package types

type RequestWithUser struct {
	User User `json:"user"`
}

type ResponseUsers struct {
	Users []User `json:"users"`
}

type UserList struct {
	Users []User `json:"users"`
}
