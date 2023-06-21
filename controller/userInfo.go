package controller

import (
	"github.com/gin-gonic/gin"
	"goPandoraAdmin-Server/services"
	"net/http"
)

// ListUsersInfo 列出所有用户信息
func ListUsersInfo(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		users := services.QueryAllUserInfo()
		c.JSON(http.StatusOK, users)
	case http.MethodPost:
		type CreateUserInfo struct {
			Password string  `form:"password" json:"password" binding:"required"`
			Email    string  `form:"email" json:"email" binding:"required,email"`
			MFA      *string `form:"mfa" json:"mfa" binding:"omitempty"`
		}

		var createUser CreateUserInfo
		err := c.ShouldBind(&createUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := services.AddUserInfo(createUser.Email, createUser.Password, *createUser.MFA)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUsersToken 更新单个用户Token
func UpdateUsersToken(c *gin.Context) {
	userID := c.Param("userID")
	switch c.Request.Method {
	case http.MethodPatch:
		info, err := services.UpdateUserInfo(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, info)
	case http.MethodDelete:
		err := services.DeleteUserInfoByUserID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	default:
		user, _ := services.QueryUserInfoByUserID(userID)
		c.JSON(http.StatusOK, user)
	}
}

func CheckAccessToken(c *gin.Context) {
	userID := c.Param("userID")
	err := services.CheckAccessToken(userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
