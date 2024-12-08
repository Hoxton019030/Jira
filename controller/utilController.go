package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpHandlerPing(c *gin.Context) {
	c.String(http.StatusOK, "收到")
}
