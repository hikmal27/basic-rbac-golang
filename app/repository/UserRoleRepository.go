package repository

import (
	"gorbac/app/entity"
	"gorbac/config"
)

type UserRoleRepository struct {
	config config.Database
}

func NewUserRoleRepository(database config.Database) UserRoleRepository {
	return UserRoleRepository{
		config: database,
	}
}

// @Summary : Find user has role
// @Description : Find user has role to database
// @Author : azrielfatur
func (r *UserRoleRepository) Find(userRole entity.UserRole) (entity.UserRole, error) {
	err := r.config.DB.Where("user_id = ? AND role_id = ?", userRole.UserID, userRole.RoleID).First(&userRole).Error

	if err != nil {
		return userRole, err
	}

	return userRole, nil
}

// @Summary : Assign user has role
// @Description : Assign user has role to database
// @Author : azrielfatur
func (r *UserRoleRepository) Attach(userRole []entity.UserRole) ([]entity.UserRole, error) {
	err := r.config.DB.Debug().Create(&userRole).Error

	if err != nil {
		return userRole, err
	}

	return userRole, nil
}

// @Summary : Un-Assign user has role
// @Description : Un-Assign user has role to database
// @Author : azrielfatur
func (r *UserRoleRepository) Detach(userRole []entity.UserRole) (bool, error) {
	userID := make([]int, 0)
	roleID := make([]int, 0)

	for _, col := range userRole {
		userID = append(userID, col.UserID)
		roleID = append(roleID, col.RoleID)
	}

	err := r.config.DB.Debug().Unscoped().Where("user_id IN ?", userID).Where("role_id IN ?", roleID).Delete(&userRole).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
