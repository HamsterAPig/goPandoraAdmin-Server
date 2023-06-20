package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goPandoraAdmin-Server/config"
	"goPandoraAdmin-Server/database"
	logger "goPandoraAdmin-Server/internal/log"
	"goPandoraAdmin-Server/model"
	"goPandoraAdmin-Server/router"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	var err error
	config.Conf, err = config.ReadConfig()
	if err != nil {
		logger.Fatal("failed to read config", zap.Error(err))
	}
	logger.InitLogger(config.Conf.DebugLevel)
	err = database.InitSQLite(config.Conf.DatabasePath)
	if err != nil {
		logger.Fatal("failed to init db", zap.Error(err))
	}
	defer database.CloseDB()
	sqlite, _ := database.GetDB()
	err = sqlite.AutoMigrate(&model.UserInfo{}, &model.AutoLoginInfo{}, &model.ShareToken{})

	if config.Conf.DebugLevel == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	adminGin := &http.Server{
		Addr:         config.Conf.Listen,
		Handler:      router.BackstageRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return adminGin.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
