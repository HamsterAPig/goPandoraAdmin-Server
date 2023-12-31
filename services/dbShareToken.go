package services

import (
	"fmt"
	"goPandoraAdmin-Server/database"
	"goPandoraAdmin-Server/internal/pandora"
	"goPandoraAdmin-Server/model"
	"time"
)

func QueryAllShareTokens() []model.ShareToken {
	db, _ := database.GetDB()
	var tokens []model.ShareToken
	db.Find(&tokens)
	return tokens
}

func QuerySingleShareToken(shareToken string) (model.ShareToken, error) {
	db, _ := database.GetDB()
	var token model.ShareToken
	res := db.Where("share_token = ?", shareToken).Find(&token)
	if res.RowsAffected == 0 {
		return token, fmt.Errorf("share token not found")
	}
	return token, nil
}

func AddShareToken(info model.CreateShareTokenRequest) (model.ShareToken, error) {
	db, _ := database.GetDB()
	var token model.ShareToken
	var user model.UserInfo
	var respond model.FakeOpenShareTokenRespond
	var fakeopen model.FakeOpenShareTokenRequest
	_, err := UpdateUserInfo(info.UserID, "")
	if err != nil {
		return token, err
	}
	res := db.Where("user_id = ?", info.UserID).Find(&user)
	if res.RowsAffected == 0 {
		return token, fmt.Errorf("user not found")
	}

	fakeopen.AccessToken = user.Token
	fakeopen.ExpiresIn = info.ExpiresTime
	fakeopen.ShowUserInfo = info.ShowUserInfo
	fakeopen.ShowConversations = info.ShowConversations
	fakeopen.UniqueName = info.UniqueName
	fakeopen.SiteLimit = info.SiteLimit
	respond, err = pandora.GetShareTokenByFakeopen(fakeopen)
	if err != nil {
		return token, fmt.Errorf("failed to get share token")
	}

	token.UserID = info.UserID
	token.UniqueName = respond.UniqueName
	token.ExpiresTime = info.ExpiresTime
	token.ExpiresTimeAt = time.Unix(respond.ExpireAt, 0)
	token.SiteLimit = &respond.SiteLimit
	token.ShareToken = respond.TokenKey
	token.ShowUserInfo = info.ShowUserInfo
	token.ShowConversations = info.ShowConversations
	token.Comment = info.Comment

	var tmpModel model.ShareToken
	res = db.Where("share_token = ?", respond.TokenKey).Find(&tmpModel)
	if res.RowsAffected == 0 {
		res = db.Create(&token)
		if res.Error != nil {
			return token, fmt.Errorf("failed to create share token in database")
		}
	} else {
		res = db.Save(token)
		if res.Error != nil {
			return token, fmt.Errorf("failed to save share token")
		}
	}

	return token, nil
}

func DeleteShareToken(fk string) error {
	db, _ := database.GetDB()
	var info model.ShareToken
	res := db.Where("share_token = ?", fk).First(&info)
	if res.RowsAffected == 0 {
		return fmt.Errorf("share token not found")
	}
	if info.ExpiresTimeAt.After(time.Now()) {
		var user model.UserInfo
		db.Where("user_id = ?", info.UserID).Find(&user)

		var fakeopen model.FakeOpenShareTokenRequest
		fakeopen.AccessToken = user.Token
		fakeopen.ExpiresIn = info.ExpiresTime
		fakeopen.ShowUserInfo = info.ShowUserInfo
		fakeopen.ShowConversations = info.ShowConversations
		fakeopen.UniqueName = info.UniqueName
		fakeopen.SiteLimit = info.SiteLimit
		_, err := pandora.GetShareTokenByFakeopen(fakeopen)
		if err != nil {
			return fmt.Errorf("failed to delete share token on fake open")
		}
	}

	res = db.Delete(&info)
	return res.Error
}

func ChangeShareTokenInfo(shareToken string, info model.ChangedShareTokenPatch) (model.ShareToken, error) {
	db, _ := database.GetDB()
	var token model.ShareToken
	res := db.Where("share_token = ?", shareToken).First(&token)
	isNeedCallUpdate := true
	if res.RowsAffected == 0 {
		return token, fmt.Errorf("share token not found")
	}
	if info.Comment != nil {
		isNeedCallUpdate = false
		token.Comment = info.Comment
	}
	if info.SiteLimit != nil {
		isNeedCallUpdate = true
		token.SiteLimit = info.SiteLimit
	}
	if info.ShowUserInfo != nil {
		isNeedCallUpdate = true
		token.ShowUserInfo = *info.ShowUserInfo
	}
	if info.ShowConversations != nil {
		isNeedCallUpdate = true
		token.ShowConversations = *info.ShowConversations
	}
	if info.ExpiresTime != nil {
		isNeedCallUpdate = true
		token.ExpiresTime = *info.ExpiresTime
	}
	db.Save(&token)
	if isNeedCallUpdate {
		return UpdateShareToken(token.ShareToken)
	}
	return token, nil
}

func UpdateShareToken(shareToken string) (model.ShareToken, error) {
	db, _ := database.GetDB()
	var token model.ShareToken
	res := db.Where("share_token = ?", shareToken).First(&token)
	if res.RowsAffected == 0 {
		return token, fmt.Errorf("share token not found")
	}
	info, err := UpdateUserInfo(token.UserID, "")
	if err != nil {
		return token, err
	}
	var createShareToken model.CreateShareTokenRequest
	createShareToken.UserID = info.UserID
	createShareToken.ExpiresTime = token.ExpiresTime
	createShareToken.ShowUserInfo = token.ShowUserInfo
	createShareToken.ShowConversations = token.ShowConversations
	createShareToken.UniqueName = token.UniqueName
	createShareToken.SiteLimit = token.SiteLimit
	createShareToken.Comment = token.Comment
	token, err = AddShareToken(createShareToken)

	return token, nil
}
