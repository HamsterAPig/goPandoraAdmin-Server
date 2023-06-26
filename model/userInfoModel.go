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
	UserID       string  `gorm:"primary_key:unique"`
	Sub          SubEnum `gorm:type:enum("google-oauth2","windowslive", "auth0") default:"auth0"`
	Token        string
	RefreshToken string
	UpdatedTime  time.Time `gorm:"autoUpdateTime"`
	ExpiryTime   time.Time
	Comment      *string
}

type CreateUserInfoRequest struct {
	Password string  `form:"password" json:"password" binding:"required"`
	Email    string  `form:"email" json:"email" binding:"required,email"`
	MFA      *string `form:"mfa" json:"mfa" binding:"omitempty"`
	Comment  *string `form:"comment" json:"comment" binding:"omitempty"`
}

type ChangeUserInfoPatch struct {
	Password *string `form:"password" json:"password" binding:"omitempty"`
	Comment  *string `form:"comment" json:"comment" binding:"omitempty"`
}
