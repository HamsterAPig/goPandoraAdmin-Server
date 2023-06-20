package model

import "time"

type SubEnum string

const (
	Google  SubEnum = "google-oauth2"
	Outlook SubEnum = "windowslive"
	OpenAI  SubEnum = "auth0"
)

type UserInfo struct {
	Email        string
	Password     string
	UserID       string  `gorm:"unique:primary_key"`
	Sub          SubEnum `gorm:type:enum("google-oauth2","windowslive", "auth0") default:"auth0"`
	Token        string
	RefreshToken string
	UpdatedTime  time.Time `gorm:"autoUpdateTime"`
	ExpiryTime   time.Time
	Comment      string
}
