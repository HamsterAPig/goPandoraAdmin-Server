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
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		req, err := services.CreateAutoLoginInfos(autoLoginInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		c.JSON(http.StatusCreated, services.RespondHandle(0, nil, req))
	default:
		infos := services.QueryAllAutoLoginInfos()
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, infos))
	}
}

func SingleAutoLoginInfosManage(c *gin.Context) {
	uuid := c.Param("UUID")
	switch c.Request.Method {
	case http.MethodDelete:
		err := services.DeleteAutoLoginInfo(uuid)
		if err != nil {
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		c.JSON(http.StatusNoContent, services.RespondHandle(0, nil, nil))
	case http.MethodPatch:
		var changeInfo model.ChangedAutoLoginInfoPatch
		err := c.ShouldBind(&changeInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		shareToken, err := services.ChangeAutoLoginInfo(uuid, changeInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, shareToken))
	default:
		UUID := c.Param("UUID")
		infos, err := services.QueryAutoLoginInfosByUUID(UUID)
		if err != nil {
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, infos))
	}
}

func UpdateAutoLoginInfo(c *gin.Context) {
	uuid := c.Param("UUID")
	info, err := services.UpdateAutoLoginInfo(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, services.RespondHandle(-1, err.Error(), nil))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, services.RespondHandle(0, nil, info))
}
