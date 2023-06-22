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

func CreateAutoLoginInfos(autoLoginInfo model.CreatedAutoLoginInfoRequest) (model.AutoLoginInfo, error) {
	db, _ := database.GetDB()
	var info model.AutoLoginInfo

	info.Comment = autoLoginInfo.Comment
	info.UserID = autoLoginInfo.UserID

	var user model.UserInfo
	res := db.Where("user_id = ?", autoLoginInfo.UserID).Find(&user)
	if res.Error != nil {
		return info, fmt.Errorf("failed to find user")
	} else if res.RowsAffected == 0 {
		return info, fmt.Errorf("user not found")
	}
	info.Token = user.Token

	res = db.Save(&info)
	if res.Error != nil {
		return info, fmt.Errorf("failed to save auto login info to db")
	}
	return info, nil
}
