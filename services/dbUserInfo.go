package services

import (
	"fmt"
	"goPandoraAdmin-Server/database"
	"goPandoraAdmin-Server/internal/pandora"
	"goPandoraAdmin-Server/model"
	"strings"
	"time"
)

// QueryAllUserInfo 查询所有用户信息
func QueryAllUserInfo() []model.UserInfo {
	db, _ := database.GetDB()
	var users []model.UserInfo

	db.Find(&users)
	return users
}

// QueryUserInfoByUserID 查询单个用户信息
func QueryUserInfoByUserID(userID string) (model.UserInfo, error) {
	db, _ := database.GetDB()
	var user model.UserInfo
	db.Where("user_id = ?", userID).Find(&user)
	return user, nil
}

func DeleteUserInfoByUserID(userID string) error {
	db, _ := database.GetDB()
	var user model.UserInfo
	res := db.Where("user_id = ?", userID).First(&user)
	if res.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	res = db.Delete(&user)
	return res.Error
}

// UpdateUserInfo 更新用户Token
func UpdateUserInfo(userID string) (model.UserInfo, error) {
	db, _ := database.GetDB()
	var user model.UserInfo
	res := db.Where("user_id = ?", userID).Find(&user)
	if res.RowsAffected == 0 {
		return user, fmt.Errorf("user not found")
	}
	token, refreshToken, err := pandora.Auth0(user.Email, user.Password, "")
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

func AddUserInfo(createInfo model.CreateUserInfoRequest) (model.UserInfo, error) {
	var user model.UserInfo
	accessToken, refreshToken, err := pandora.Auth0(createInfo.Email, createInfo.Password, *createInfo.MFA)
	if err != nil {
		return user, fmt.Errorf("auth0 error: %s", err)
	}
	dec, err := pandora.CheckAccessToken(accessToken)
	if err != nil {
		return model.UserInfo{}, fmt.Errorf("check access token error: %s", err)
	}

	user.Email = createInfo.Email
	user.Password = createInfo.Password
	user.Token = accessToken
	user.RefreshToken = refreshToken
	user.Sub = model.SubEnum(strings.Split(dec.Sub, "|")[0])
	user.UserID = dec.Auth.UserID
	user.ExpiryTime = time.Unix(int64(dec.Exp), 0)
	user.Comment = *createInfo.Comment

	info, err := updateUserInfo(user)
	if err != nil {
		return info, err
	}
	return user, nil
}

// updateUserInfo 更新用户信息表
func updateUserInfo(user model.UserInfo) (model.UserInfo, error) {
	db, _ := database.GetDB()
	res := db.Save(&user)
	if res.Error != nil {
		return user, fmt.Errorf("create user error: %s", res.Error)
	}
	return user, nil
}

func CheckAccessToken(userID string) error {
	var user model.UserInfo
	db, _ := database.GetDB()
	res := db.Where("user_id = ?", userID).Find(&user)
	if res.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	_, err := pandora.CheckAccessToken(user.Token)
	if err != nil {
		return fmt.Errorf("check access token error: %s", err)
	}
	return nil
}
