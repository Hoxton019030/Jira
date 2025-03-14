package service

import (
	"jira/repository"
	"jira/request"
	"log"
)

func AddUser(addUserRequest request.AddUserRequest) error {
	user := addUserRequest.ConvertToUser()
	err := repository.AddUser(user)
	if err != nil {
		log.Fatal("創建user失敗 ", err)
	}
	return err
}
