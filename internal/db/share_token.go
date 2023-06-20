package db

import (
	"time"
)

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

type faseOpenShareToken struct {
	ExpireAt          int64  `json:"expire_at"`
	ShowConversations bool   `json:"show_conversations"`
	ShowUserinfo      bool   `json:"show_userinfo"`
	SiteLimit         string `json:"site_limit"`
	TokenKey          string `json:"token_key"`
	UniqueName        string `json:"unique_name"`
}
