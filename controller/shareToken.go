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
			return
		}
		token, err := services.AddShareToken(shareTokenInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, services.RespondHandle(-1, err.Error(), nil))
			return
		}
		c.JSON(http.StatusCreated, services.RespondHandle(0, nil, token))
	default:
		sk := services.QueryAllShareTokens()
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, sk))
	}
}

func SingleShareTokenManage(c *gin.Context) {
	id := c.Param("id")
	switch c.Request.Method {
	case http.MethodDelete:
		err := services.DeleteShareToken(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, services.RespondHandle(-1, err.Error(), nil))
			return
		}
		c.JSON(http.StatusNoContent, services.RespondHandle(0, nil, nil))
	default:
		id := c.Param("id")
		sk := services.QuerySingleShareToken(id)
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, sk))
	}
}
