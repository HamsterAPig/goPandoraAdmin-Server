package services

import (
	"goPandoraAdmin-Server/database"
	"goPandoraAdmin-Server/model"
)

// QueryAllAutoLoginInfos 查询所有自动登录信息
func QueryAllAutoLoginInfos() []model.AutoLoginInfo {
	db, _ := database.GetDB()
	var autoLoginInfo []model.AutoLoginInfo
	db.Find(&autoLoginInfo)
	return autoLoginInfo
}
