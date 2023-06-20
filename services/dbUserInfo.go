package services

import (
	"fmt"
	"goPandoraAdmin-Server/database"
	"goPandoraAdmin-Server/internal/pandora"
	"goPandoraAdmin-Server/model"
)

// QueryAllUserInfo 查询所有用户信息
func QueryAllUserInfo() []model.UserInfo {
	db, _ := database.GetDB()
	var users []model.UserInfo

	db.Find(&users)
	return users
}

// UpdateUserInfo 更新用户Token
func UpdateUserInfo(userID string) (model.UserInfo, error) {
	db, _ := database.GetDB()
	var user model.UserInfo
	res := db.Where("user_id = ?", userID).Find(&user)
	if res.RowsAffected == 0 {
		return user, fmt.Errorf("user not found")
	}
	token, refreshToken, err := pandora.Auth0(user.Email, user.Password, "", "")
	if err != nil {
		return user, fmt.Errorf("auth0 error: %s", err)
	}
	user.Token = token
	user.RefreshToken = refreshToken
	_, err = pandora.CheckAccessToken(user.Token)
	if err != nil {
		return user, fmt.Errorf("check access token error: %s", err)
	}
	db.Save(&user)
	return user, nil
}
