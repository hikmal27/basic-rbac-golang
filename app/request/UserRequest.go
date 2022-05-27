package request

import "gorbac/app/entity"

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	Username  string `json:"username" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type UserRole struct {
	UserID    int    `json:"user_id"`
	RoleID    int    `json:"role_id"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type UserHasRole struct {
	Roles []entity.UserRole `json:"roles"`
}
