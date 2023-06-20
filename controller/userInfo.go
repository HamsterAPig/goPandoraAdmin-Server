package controller

import (
	"github.com/gin-gonic/gin"
	"goPandoraAdmin-Server/services"
	"net/http"
)

// ListUsersInfo 列出所有用户信息
func ListUsersInfo(c *gin.Context) {
	users := services.QueryAllUserInfo()
	c.JSON(http.StatusOK, users)
}
