package entity

import (
	"time"

	"gorm.io/gorm"
)

type RolePermission struct {
	ID           int `gorm:"primary_key; not null"`
	RoleID       int `gorm:"not null; index" json:"role_id"`
	PermissionID int `gorm:"not null; index" json:"permission_id"`
	CreatedAt    time.Time
	CreatedBy    string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt    time.Time
	UpdatedBy    string `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt    gorm.DeletedAt
}

type TablerRolePermission interface {
	TableName() string
}

func (RolePermission) TableName() string {
	return "tr_role_permissions"
}
