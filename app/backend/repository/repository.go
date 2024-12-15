package repository

import (
	"errors"
	"src/initializers"
	"src/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	initializers.DB.First(&user, "email = ?", email)
	if user.UserId == 0 {
		return nil, errors.New("invalid email")
	}
	return &user, nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := initializers.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("invalid id")
	}
	return &user, nil
}
