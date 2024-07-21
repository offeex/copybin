package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username     string `gorm:"uniqueIndex;not null"`
	Role         string
	Pastas       []*Pasta       `gorm:"foreignKey:UserID"`
	Integrations []*Integration `gorm:"foreignKey:UserID"`
}

type Pasta struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Text   string
	UserID uint
}

type Integration struct {
	*gorm.Model
	Provider     string `gorm:"uniqueIndex:idx_user_provider"`
	ProviderID   string
	AccessToken  string
	RefreshToken string
	UserID       uint `gorm:"uniqueIndex:idx_user_provider"`
}
