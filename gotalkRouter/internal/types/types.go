package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type User struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
}

type RequestWithUser struct {
	User User `json:"user"`
}

type ResponseWithMessage struct {
	Message string `json:"message"`
}

type ResponseUsers struct {
	Users []User `json:"users"`
}
