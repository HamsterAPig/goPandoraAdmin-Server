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

func UpdateUsersToken(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPut:
		userID := c.Param("userID")
		info, err := services.UpdateUserInfo(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, info)
	}
}
