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
		c.JSON(http.StatusOK, sk)
	}
}

func SingleShareTokenManage(c *gin.Context) {

}
