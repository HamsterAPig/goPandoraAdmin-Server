package model

import "time"

type ShareToken struct {
	ID            uint `gorm:"primary_key:autoIncrement"`
	UserID        string
	UniqueName    string
	ExpiresTime   int64
	ExpiresTimeAt time.Time
	SiteLimit     string
	SK            string    `gorm:"unique"`
	UpdateTime    time.Time `gorm:"autoUpdateTime"`
	Comment       string
}

type fakeOpenShareToken struct {
	ExpireAt          int64  `json:"expire_at"`
	ShowConversations bool   `json:"show_conversations"`
	ShowUserinfo      bool   `json:"show_userinfo"`
	SiteLimit         string `json:"site_limit"`
	TokenKey          string `json:"token_key"`
	UniqueName        string `json:"unique_name"`
}

type CreateShareTokenRequest struct {
	UserID      string  `form:"user-id" json:"user-id" binding:"required"`
	UniqueName  string  `form:"unique-name" json:"unique-name" binding:"required"`
	ExpiresTime int64   `form:"expires-time" json:"expires-time" binding:"required"`
	SiteLimit   *string `form:"site-limit" json:"site-limit"`
	Comment     *string `form:"comment" json:"comment"`
}
