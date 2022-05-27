package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gprm:"primary_key; not null"`
	Username  string `gorm:"size:100; not null"`
	Name      string `gorm:"size:100; not null"`
	Email     string `gorm:"size:100; not null"`
	Password  string `gorm:"size:100; not null"`
	Roles     []Role `gorm:"many2many:tr_user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	CreatedBy string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt time.Time
	UpdatedBy string `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt gorm.DeletedAt
}

type TablerUser interface {
	TableName() string
}

func (User) TableName() string {
	return "ms_users"
}
