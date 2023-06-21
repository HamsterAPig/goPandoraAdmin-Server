package model

type CreateUserInfoRequest struct {
	Password string  `form:"password" json:"password" binding:"required"`
	Email    string  `form:"email" json:"email" binding:"required,email"`
	MFA      *string `form:"mfa" json:"mfa" binding:"omitempty"`
	Comment  *string `form:"comment" json:"comment" binding:"omitempty"`
}
