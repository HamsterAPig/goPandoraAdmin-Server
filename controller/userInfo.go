package controller

import (
	"github.com/gin-gonic/gin"
	"goPandoraAdmin-Server/model"
	"goPandoraAdmin-Server/services"
	"net/http"
)

// UserInfosManage 列出所有用户信息
func UserInfosManage(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		users := services.QueryAllUserInfo()
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, users))
	case http.MethodPost:
		var createUser model.CreateUserInfoRequest
		err := c.ShouldBind(&createUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
		}
		user, err := services.AddUserInfo(createUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, services.RespondHandle(-1, err.Error(), nil))
		}
		c.JSON(http.StatusCreated, services.RespondHandle(0, nil, user))
	}
}

// SingleUserInfosManage 更新单个用户Token
func SingleUserInfosManage(c *gin.Context) {
	userID := c.Param("userID")
	switch c.Request.Method {
	case http.MethodPatch:
		info, err := services.UpdateUserInfo(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, services.RespondHandle(-1, err.Error(), nil))
		}
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, info))
	case http.MethodDelete:
		err := services.DeleteUserInfoByUserID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, services.RespondHandle(-1, err.Error(), nil))
		}
		c.JSON(http.StatusNoContent, services.RespondHandle(0, nil, nil))
	default:
		user, _ := services.QueryUserInfoByUserID(userID)
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, user))
	}
}

func CheckAccessToken(c *gin.Context) {
	userID := c.Param("userID")
	err := services.CheckAccessToken(userID)
	if err != nil {
		c.JSON(http.StatusOK, services.RespondHandle(-1, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, services.RespondHandle(0, "healthy", nil))
}
