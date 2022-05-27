package routes

import (
	"gorbac/app/controller"
	"gorbac/app/repository"
	"gorbac/app/service"
	"gorbac/config"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func WebRouter(db config.Database) {
	// Repository Asset
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)
	roleMenuRepo := repository.NewRoleMenuRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)
	rolePermissionRepo := repository.NewRolePermissionRepository(db)

	// Service Asset
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	roleService := service.NewRoleService(roleRepo)
	menuService := service.NewMenuService(menuRepo)
	userRoleService := service.NewUserRoleService(userRoleRepo)
	roleMenuService := service.NewRoleMenuService(roleMenuRepo)
	permissionService := service.NewPermissionService(permissionRepo)
	rolePermissionService := service.NewRolePermissionService(rolePermissionRepo)

	//Controller Asset
	userController := controller.NewUserController(userService, userRoleService)
	authController := controller.NewAuthController(userService, authService)
	roleController := controller.NewRoleController(roleService, roleMenuService, rolePermissionService)
	menuController := controller.NewMenuController(menuService)
	permController := controller.NewPermissionController(permissionService)

	// Route
	httpRouter := gin.Default()

	// Register routing
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Testing  connection
	httpRouter.GET("/status-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Service ✅ API Up and Running"})
	})

	v1 := httpRouter.Group("/api/v1") // Grouping routes

	// Auth routes
	v1.POST("/auth/register", authController.Register)
	v1.POST("/auth/login", authController.Login)

	// Users routes
	v1.GET("/users", userController.Index)
	v1.GET("/users/:id", userController.Show)
	v1.PUT("/users/:id", userController.Update)
	v1.DELETE("/users/:id", userController.Delete)

	// User role routes
	v1.POST("/users/roles/assign", userController.AttachRole)
	v1.POST("/users/roles/unassign", userController.DetachRole)

	// Roles routes
	v1.POST("/roles", roleController.Store)
	v1.GET("/roles", roleController.Index)
	v1.GET("/roles/:id", roleController.Show)
	v1.PUT("/roles/:id", roleController.Update)
	v1.DELETE("/roles/:id", roleController.Delete)

	// Role menu routes
	v1.POST("/roles/menus", roleController.AttachMenu)
	v1.DELETE("/roles/menus", roleController.DetachMenu)

	// Roles routes
	v1.POST("/permissions", permController.Store)
	v1.GET("/permissions", permController.Index)
	v1.GET("/permissions/:id", permController.Show)
	v1.PUT("/permissions/:id", permController.Update)
	v1.DELETE("/permissions/:id", permController.Delete)

	// Role permission routes
	v1.POST("/roles/permissions", roleController.AttachPermission)
	v1.DELETE("/roles/permissions", roleController.DetachPermission)

	// Menus routes
	v1.POST("/menus", menuController.Store)
	v1.GET("/menus", menuController.Index)
	v1.GET("/menus/:id", menuController.Show)
	v1.PUT("/menus/:id", menuController.Update)
	v1.DELETE("/menus/:id", menuController.Delete)

	httpRouter.Run(":" + os.Getenv("APP_PORT")) // Routes
}
