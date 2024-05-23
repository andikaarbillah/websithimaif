package repository

import (
	"websitehimaif/model"

	"gorm.io/gorm"
)

type UserRepository interface{
	Register(user model.User)(*model.User, error)
}

type userRepository struct{
	db *gorm.DB
}

func(ur *userRepository)Register(user model.User)(*model.User, error){
	result:= ur.db.Create(&user)
	if result.Error != nil{
		return nil, result.Error
	}

	return &user, nil
}

func NewUserRepository(db *gorm.DB)UserRepository{
	return &userRepository{db: db}
}