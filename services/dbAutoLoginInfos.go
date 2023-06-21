package services

import (
	"fmt"
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

// QueryAllAutoLoginInfosByUUID 查询单个自动登录信息
func QueryAllAutoLoginInfosByUUID(UUID string) (model.AutoLoginInfo, error) {
	db, _ := database.GetDB()
	var autoLoginInfo model.AutoLoginInfo
	res := db.Where("uuid = ?", UUID).Find(&autoLoginInfo)
	if res.Error != nil {
		return autoLoginInfo, res.Error
	}
	if res.RowsAffected == 0 {
		return autoLoginInfo, fmt.Errorf("user not found")
	}
	return autoLoginInfo, nil
}
