package request

import (
	"jira/model"
	"log"
)

type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (addUserRequest AddUserRequest) ConvertToUser() model.User {
	log.Print("轉換型態")
	return model.User{
		Username: addUserRequest.Username,
		Password: addUserRequest.Password,
	}
}
