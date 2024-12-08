package router

import (
	"github.com/gin-gonic/gin"
	"jira/controller"
	"log"
)

var router *gin.Engine

func Init() {
	router = gin.New()
	router.GET("/ping", controller.HttpHandlerPing)
}

func Run() {
	err := router.Run()
	if err != nil {
		log.Fatal("Run發生錯誤...")
		return
	}
}
