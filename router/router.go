package router

import (
	"websitehimaif/controller"
	"websitehimaif/repository"
	"websitehimaif/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Router(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserServices(userRepo)
	userController := controller.NewUserController(userService, ctx)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/register", userController.VerifikasiRegister)
	}
}
