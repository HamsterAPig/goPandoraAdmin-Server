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
		v1.Any("/users", controller.UserInfosManage)
		v1.Any("/users/:userID", controller.SingleUserInfosManage)
		v1.GET("/users/:userID/check", controller.CheckAccessToken)
		v1.GET("/users/:userID/update", controller.UpdateAccessToken)

		v1.Any("/auto-login-infos", controller.AutoLoginInfosManage)
		v1.Any("/auto-login-infos/:UUID", controller.SingleAutoLoginInfosManage)

		v1.Any("/share-tokens", controller.ShareTokensManage)
		v1.Any("/share-tokens/:id", controller.SingleShareTokenManage)
	}
	return r
}
