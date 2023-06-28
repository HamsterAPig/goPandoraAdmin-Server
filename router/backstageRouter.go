package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"goPandoraAdmin-Server/config"
	"goPandoraAdmin-Server/controller"
	logger "goPandoraAdmin-Server/internal/log"
	"net/http"
)

func BackstageRouter() http.Handler {
	r := gin.Default()
	if config.Conf.AllowCors {
		r.Use(cors.Default())
	}
	APIPath := "/api/v1"
	if config.Conf.EnableUUIDURI {
		APIPath = uuid.New().String() + "/" + APIPath
		logger.Info("Runing path is ", zap.String("APIPath", APIPath))
	}
	v1 := r.Group(APIPath)
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
	return r
}
