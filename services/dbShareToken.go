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

func QuerySingleShareToken(id string) model.ShareToken {
	db, _ := database.GetDB()
	var token model.ShareToken
	db.Where("id = ?", id).Find(&token)
	return token
}

func AddShareToken(info model.CreateShareTokenRequest) (model.ShareToken, error) {
	db, _ := database.GetDB()
	var token model.ShareToken
	var user model.UserInfo
	res := db.Where("user_id = ?", info.UserID).Find(&user)
	if res.RowsAffected == 0 {
		return token, fmt.Errorf("user not found")
	}

	var fakeopen model.FakeOpenShareTokenRequest
	fakeopen.AccessToken = user.Token
	fakeopen.ExpiresIn = info.ExpiresTime
	fakeopen.ShowUserInfo = info.ShowUserInfo
	fakeopen.ShowConversations = info.ShowConversations
	fakeopen.UniqueName = info.UniqueName
	fakeopen.SiteLimit = info.SiteLimit
	respond, err := pandora.GetShareTokenByFakeopen(fakeopen)
	if err != nil {
		return token, fmt.Errorf("failed to get share token")
	}

	token.UserID = info.UserID
	token.UniqueName = respond.UniqueName
	token.ExpiresTime = info.ExpiresTime
	token.ExpiresTimeAt = time.Unix(respond.ExpireAt, 0)
	token.SiteLimit = &respond.SiteLimit
	token.SK = respond.TokenKey
	token.ShowUserInfo = info.ShowUserInfo
	token.ShowConversations = info.ShowConversations

	res = db.Save(&token)
	if res.Error != nil {
		return token, fmt.Errorf("failed to save share token")
	}

	return token, nil
}

func DeleteShareToken(id string) error {
	db, _ := database.GetDB()
	var info model.ShareToken
	res := db.Where("id = ?", id).First(&info)
	if res.RowsAffected == 0 {
		return fmt.Errorf("share token not found")
	}
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

	res = db.Delete(&info)
	return res.Error
}

func ChangeShareTokenInfo(id string, info model.ChangedShareTokenPatch) (model.ShareToken, error) {
	db, _ := database.GetDB()
	var token model.ShareToken
	res := db.Where("id = ?", id).First(&token)
	if res.RowsAffected == 0 {
		return token, fmt.Errorf("share token not found")
	}
	if info.Comment != nil {
		token.Comment = info.Comment
	}
	if info.SiteLimit != nil {
		token.SiteLimit = info.SiteLimit
	}
	if info.ShowUserInfo != nil {
		token.ShowUserInfo = *info.ShowUserInfo
	}
	if info.ShowConversations != nil {
		token.ShowConversations = *info.ShowConversations
	}
	if info.ExpiresTime != nil {
		token.ExpiresTime = *info.ExpiresTime
	}
	db.Save(&token)
	return token, nil
}
