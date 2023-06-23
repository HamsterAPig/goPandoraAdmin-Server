package controller

import (
	"github.com/gin-gonic/gin"
	"goPandoraAdmin-Server/services"
	"net/http"
)

func ShareTokensManage(c *gin.Context) {
	switch c.Request.Method {
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
