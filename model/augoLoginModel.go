package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AutoLoginInfo struct {
	UUID    uuid.UUID `gorm:"primaryKey;type:char(36);not null;unique"`
	UserID  string
	Token   string
	Comment string
}

// BeforeCreate 向User表插入数据后自动添加UUID
func (u *AutoLoginInfo) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.New()
	return nil
}
