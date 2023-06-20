package main

import (
	"go.uber.org/zap"
	"goPandoraAdmin-Server/config"
	"goPandoraAdmin-Server/database"
	logger "goPandoraAdmin-Server/internal/log"
	"goPandoraAdmin-Server/model"
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
}
