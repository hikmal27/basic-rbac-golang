package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	ID        int `gprm:"primary_key; not null"`
	UserID    int `gorm:"not null; index" json:"user_id"`
	RoleID    int `gorm:"not null; index" json:"role_id"`
	CreatedAt time.Time
	CreatedBy string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt time.Time
	UpdatedBy string `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt gorm.DeletedAt
}

type TablerUserRole interface {
	TableName() string
}

func (UserRole) TableName() string {
	return "tr_user_roles"
}
