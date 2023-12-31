package services

import (
	"fmt"
	"goPandoraAdmin-Server/database"
	logger "goPandoraAdmin-Server/internal/log"
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
func UpdateUserInfo(userID string, forceT string) (model.UserInfo, error) {
	var user model.UserInfo
	var force bool
	if forceT == "true" || forceT == "True" {
		force = true
	} else {
		force = false
	}
	db, _ := database.GetDB()
	res := db.Where("user_id = ?", userID).Find(&user)
	if res.RowsAffected == 0 {
		return user, fmt.Errorf("user not found")
	}

	needUpdate := user.ExpiryTime.Before(time.Now())
	if needUpdate || force {
		logger.Info("begin update token by refresh token")
		token, err := pandora.GetTokenByRefreshToken(user.RefreshToken)
		if err != nil {
			return user, fmt.Errorf("refresh token error: %s", err)
		}
		user.Token = token
		payload, err := pandora.CheckAccessToken(user.Token)
		if err != nil {
			return user, fmt.Errorf("check access token error: %s", err)
		}
		user.ExpiryTime = time.Unix(int64(payload.Exp), 0)
		db.Save(&user)
	}
	return user, nil
}

func AddUserInfo(createInfo model.CreateUserInfoRequest) (model.UserInfo, error) {
	var user model.UserInfo
	var mfaCode = ""
	if createInfo.MFA != nil {
		mfaCode = *createInfo.MFA
	}
	accessToken, refreshToken, err := pandora.Auth0(createInfo.Email, createInfo.Password, mfaCode)
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
	user.Comment = createInfo.Comment

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

func ChangeUserInfo(userID string, changeInfo model.ChangeUserInfoPatch) (model.UserInfo, error) {
	var user model.UserInfo
	db, _ := database.GetDB()
	res := db.Where("user_id = ?", userID).Find(&user)
	if res.RowsAffected == 0 {
		return user, fmt.Errorf("user not found")
	}
	if changeInfo.Password != nil {
		user.Password = *changeInfo.Password
	}
	if changeInfo.Comment != nil {
		user.Comment = changeInfo.Comment
	}
	db.Save(&user)
	return user, nil
}
