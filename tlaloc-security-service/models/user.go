package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	Role      string `gorm:"default:user" json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  bool   `gorm:"default:true" json:"is_active"`
}

type RefreshToken struct {
	gorm.Model
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	Token     string    `gorm:"uniqueIndex;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Revoked   bool      `gorm:"default:false" json:"revoked"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
}

type UserSession struct {
	gorm.Model
	UserID     uint      `gorm:"index;not null" json:"user_id"`
	SessionID  string    `gorm:"uniqueIndex;not null" json:"session_id"`
	DeviceInfo string    `json:"device_info"`
	IPAddress  string    `json:"ip_address"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`
	ExpiresAt  time.Time `gorm:"not null" json:"expires_at"`
	LastActive time.Time `json:"last_active"`
	Revoked    bool      `gorm:"default:false" json:"revoked"`
}
