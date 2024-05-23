package controller

import (
	"websitehimaif/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us  services.UserService
	ctx *gin.Context
}

func (uc UserController) VerifikasiRegister(ctx *gin.Context) {
	data, err := uc.us.Register(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "error",
			"data":    data,
		})
	}
	ctx.JSON(200, gin.H{
		"message": "successfully",
		"data":    data,
	})
}

func NewUserController(us services.UserService, ctx *gin.Context) UserController {
	return UserController{
		us:  us,
		ctx: ctx,
	}
}
