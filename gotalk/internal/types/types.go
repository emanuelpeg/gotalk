package types

type User struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
}

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type RequestWithUserName struct {
	UserName string `json:"username"`
}

type ResponseWithSessionId struct {
	Id string `json:"id"`
}

type RequestWithSessionId struct {
	Id string `json:"id"`
}
