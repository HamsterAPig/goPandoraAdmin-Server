package controller

import (
	"github.com/gin-gonic/gin"
	"goPandoraAdmin-Server/database"
	"goPandoraAdmin-Server/model"
	"net/http"
)

// ListUsersInfo 列出所有用户信息
func ListUsersInfo(c *gin.Context) {
	db, _ := database.GetDB()
	var users []model.UserInfo

	db.Find(&users)
	c.JSON(http.StatusOK, users)
}
