package model

import "github.com/google/uuid"

type AutoLoginInfo struct {
	UUID    uuid.UUID `gorm:"primaryKey;type:char(36);not null;unique"`
	UserID  string
	Token   string
	Comment string
}
