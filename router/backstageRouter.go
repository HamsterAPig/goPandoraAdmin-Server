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
		v1.Any("/users", controller.ListUsersInfo)
		v1.Any("/users/:userID/token", controller.UpdateUsersToken)
		v1.GET("/users/:userID/check", controller.CheckAccessToken)
	}
	return r
}
