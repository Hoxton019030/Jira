package controller

import (
	"github.com/gin-gonic/gin"
	"jira/request"
	"jira/service"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var addUserRequest request.AddUserRequest
	err := c.ShouldBindJSON(&addUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.AddUser(addUserRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, addUserRequest)

}
