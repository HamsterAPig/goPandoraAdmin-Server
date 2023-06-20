package services

import (
	"goPandoraAdmin-Server/database"
	"goPandoraAdmin-Server/model"
)

// QueryAllUserInfo 查询所有用户信息
func QueryAllUserInfo() []model.UserInfo {
	db, _ := database.GetDB()
	var users []model.UserInfo

	db.Find(&users)
	return users
}
