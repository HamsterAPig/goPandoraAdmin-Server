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
		}
		token, err := services.AddShareToken(shareTokenInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, services.RespondHandle(-1, err.Error(), nil))
		}
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, token))
	default:
		sk := services.QueryAllShareTokens()
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, sk))
	}
}

func SingleShareTokenManage(c *gin.Context) {
	switch c.Request.Method {
	default:
		id := c.Param("id")
		sk := services.QuerySingleShareToken(id)
		c.JSON(http.StatusOK, services.RespondHandle(0, nil, sk))
	}
}
