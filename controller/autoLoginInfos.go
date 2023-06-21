package controller

import (
	"github.com/gin-gonic/gin"
	"goPandoraAdmin-Server/services"
	"net/http"
)

func AutoLoginInfos(c *gin.Context) {
	switch c.Request.Method {
	default:
		infos := services.QueryAllAutoLoginInfos()
		c.JSON(http.StatusOK, infos)
	}
}
