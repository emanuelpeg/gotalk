package types

import (
	"sync"
)

var lock = &sync.Mutex{}

type UserList struct {
	Users []User `json:"users"`
}

var userList *UserList

func GetUserList() *UserList {
	if userList == nil {
		lock.Lock()
		defer lock.Unlock()
		if userList == nil {
			userList = &UserList{Users: make([]User, 0)}
		}
	}

	return userList
}

func AddUser(user User) {
	var users = GetUserList().Users
	users = append(users, user)
	userList.Users = users
}

func DeleteUser(user User) bool {
	var users = GetUserList().Users
	var index = -1
	for i, anUser := range users {
		if anUser == user {
			index = i
			break
		}
	}
	if index != -1 {
		users = append(users[:index], users[index+1:]...)
		userList.Users = users
		return true
	}
	return false
}
