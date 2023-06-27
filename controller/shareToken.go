package controller

import (
	"github.com/gin-gonic/gin"
	"goPandoraAdmin-Server/model"
	"goPandoraAdmin-Server/services"
	"net/http"
)

func ShareTokensManage(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		var shareTokenInfo model.CreateShareTokenRequest
		err := c.ShouldBind(&shareTokenInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		token, err := services.AddShareToken(shareTokenInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		c.JSON(http.StatusCreated, services.RespondHandle(0, nil, token))
	default:
		sk := services.QueryAllShareTokens()
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, sk))
	}
}

func SingleShareTokenManage(c *gin.Context) {
	fk := c.Param("fk")
	switch c.Request.Method {
	case http.MethodDelete:
		err := services.DeleteShareToken(fk)
		if err != nil {
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		c.JSON(http.StatusNoContent, services.RespondHandle(0, nil, nil))
	case http.MethodPatch:
		var changeInfo model.ChangedShareTokenPatch
		err := c.ShouldBind(&changeInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		shareToken, err := services.ChangeShareTokenInfo(fk, changeInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, shareToken))
	default:
		sk, err := services.QuerySingleShareToken(fk)
		if err != nil {
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, sk))
	}
}

func UpdateShareToken(c *gin.Context) {
	fk := c.Param("fk")
	info, err := services.UpdateShareToken(fk)
	if err != nil {
		c.JSON(http.StatusInternalServerError, services.RespondHandle(-1, err.Error(), nil))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, services.RespondHandle(0, nil, info))
}
