package router

import (
	"github.com/gin-gonic/gin"
	"goPandoraAdmin-Server/controller"
	"net/http"
)

func BackstageRouter() http.Handler {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", controller.ListUsersInfo)
	}
	return r
}
