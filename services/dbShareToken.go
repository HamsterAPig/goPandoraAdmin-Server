package services

import (
	"goPandoraAdmin-Server/database"
	"goPandoraAdmin-Server/model"
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
