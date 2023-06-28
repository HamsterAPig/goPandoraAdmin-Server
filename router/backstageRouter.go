package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goPandoraAdmin-Server/config"
	"goPandoraAdmin-Server/controller"
	"net/http"
)

func BackstageRouter() http.Handler {
	r := gin.Default()
	if config.Conf.AllowCors {
		r.Use(cors.Default())
	}
	v1 := r.Group(config.Conf.EnableUUIDURI + "/api/v1")
	{
		v1.Any("/users", controller.UserInfosManage)
		v1.Any("/users/:userID", controller.SingleUserInfosManage)
		v1.GET("/users/:userID/check", controller.CheckAccessToken)
		v1.GET("/users/:userID/update", controller.UpdateAccessToken)

		v1.Any("/auto-login-infos", controller.AutoLoginInfosManage)
		v1.Any("/auto-login-infos/:UUID", controller.SingleAutoLoginInfosManage)
		v1.GET("/auto-login-infos/:UUID/update", controller.UpdateAutoLoginInfo)

		v1.Any("/share-tokens", controller.ShareTokensManage)
		v1.Any("/share-tokens/:fk", controller.SingleShareTokenManage)
		v1.GET("/share-tokens/:fk/update", controller.UpdateShareToken)
	}
	r.GET("/auto-login-infos/:UUID", controller.SingleAutoLoginInfosManage)
	return r
}
