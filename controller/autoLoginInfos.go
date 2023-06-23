package controller

import (
	"github.com/gin-gonic/gin"
	"goPandoraAdmin-Server/model"
	"goPandoraAdmin-Server/services"
	"net/http"
)

// AutoLoginInfosManage 自动登陆信息管理
func AutoLoginInfosManage(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		var autoLoginInfo model.CreatedAutoLoginInfoRequest
		if err := c.ShouldBind(&autoLoginInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		req, err := services.CreateAutoLoginInfos(autoLoginInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, req)
	default:
		infos := services.QueryAllAutoLoginInfos()
		c.JSON(http.StatusOK, infos)
	}
}

func SingleAutoLoginInfosManage(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodDelete:
		uuid := c.Param("UUID")
		err := services.DeleteAutoLoginInfo(uuid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{
			"success": true,
		})
	default:
		UUID := c.Param("UUID")
		infos, err := services.QueryAllAutoLoginInfosByUUID(UUID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, infos)
	}
}
