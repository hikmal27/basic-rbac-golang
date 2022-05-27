package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          int          `gprm:"primary_key; not null"`
	Name        string       `gorm:"size:100; not null"`
	Description string       `gorm:"default: NULL"`
	Menus       []Menu       `gorm:"many2many:tr_role_menus;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Permissions []Permission `gorm:"many2many:tr_role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time
	CreatedBy   string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt   time.Time
	UpdatedBy   string `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt   gorm.DeletedAt
}

type TablerRole interface {
	TableName() string
}

func (Role) TableName() string {
	return "ms_roles"
}
