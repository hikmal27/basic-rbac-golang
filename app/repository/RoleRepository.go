package repository

import (
	"gorbac/app/entity"
	"gorbac/config"
)

type RoleRepository struct {
	config config.Database
}

func NewRoleRepository(database config.Database) RoleRepository {
	return RoleRepository{
		config: database,
	}
}

// @Summary : Insert role
// @Description : Insert role to database
// @Author : azrielfatur
func (r *RoleRepository) Insert(role entity.Role) (entity.Role, error) {
	err := r.config.DB.Create(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

// @Summary : Get roles
// @Description : -
// @Author : azrielfatur
func (r *RoleRepository) FindAll() ([]entity.Role, error) {
	var roles []entity.Role

	err := r.config.DB.Find(&roles).Error

	if err != nil {
		return roles, err
	}

	return roles, nil
}

// @Summary : Get role
// @Description : Find role by ID
// @Author : azrielfatur
func (r *RoleRepository) FindById(ID int) (entity.Role, error) {
	var user entity.Role

	err := r.config.DB.Where("id = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Update role
// @Description : Update role by ID
// @Author : azrielfatur
func (r *RoleRepository) Update(role entity.Role, ID int) (entity.Role, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

// @Summary : Delete role
// @Description : Delete role temporary
// @Author : azrielfatur
func (r *RoleRepository) Delete(ID int) (bool, error) {
	var role entity.Role

	err := r.config.DB.Where("id = ?", ID).Delete(&role).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
