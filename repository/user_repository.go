package repository

import (
	"jira/database"
	"jira/model"
	"log"
)

func AddUser(user model.User) error {
	result := database.DB.Create(&user)
	log.Printf("創建user成功 %s", result.Error)
	return result.Error
}
