package services

import (
	"websitehimaif/hashing"
	"websitehimaif/id"
	"websitehimaif/model"
	"websitehimaif/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	Register(ctx *gin.Context) (*model.User, error)
}

type userService struct {
	ur repository.UserRepository
}

func (us *userService) Register(ctx *gin.Context) (*model.User, error) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		return nil, err
	}

	usr := model.User{
		ID:        id.IdRndm(8),
		Username:  user.Username,
		Email:     user.Email,
		Password:  hashing.Hashing(user.Password),
		Address:   user.Address,
		Photo:     user.Photo,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		IsDeleted: user.IsDeleted,
	}

	result, err := us.ur.Register(usr)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func NewUserServices(ur repository.UserRepository) UserService {
	return &userService{
		ur: ur,
	}
}
