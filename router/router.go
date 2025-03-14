package router

import (
	"github.com/gin-gonic/gin"
	"jira/controller"
	"log"
)

var router *gin.Engine

func Init() {
	log.Println("初始化路由")
	router = gin.New()
	router.GET("/ping", controller.HttpHandlerPing)
	router.POST("/user", controller.CreateUser)
}

func Run() {
	err := router.Run()
	if err != nil {
		log.Fatal("Run發生錯誤...")
		return
	}
}
