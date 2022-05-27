package entity

import (
	"time"

	"gorm.io/gorm"
)

type RoleMenu struct {
	ID        int `gorm:"primary_key; not null"`
	RoleID    int `gorm:"not null; index" json:"role_id"`
	MenuID    int `gorm:"not null; index" json:"menu_id"`
	CreatedAt time.Time
	CreatedBy string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt time.Time
	UpdatedBy string `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt gorm.DeletedAt
}

type TablerRoleMenu interface {
	TableName() string
}

func (RoleMenu) TableName() string {
	return "tr_role_menus"
}
