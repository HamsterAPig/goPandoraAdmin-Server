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
