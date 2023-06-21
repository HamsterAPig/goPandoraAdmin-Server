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

func AutoLoginInfosByUUID(c *gin.Context) {
	switch c.Request.Method {
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
