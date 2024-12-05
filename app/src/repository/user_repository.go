package repository

import (
	"src/initializers"
	"src/models"
)

// CreateUser Функция для создания нового пользователя
func CreateUser(user *models.User) error {
	// Создаем нового пользователя в базе данных
	result := initializers.DB.Create(&user)
	return result.Error
}

// FindUserByEmail Функция для поиска пользователя по email
func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := initializers.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser Функция для обновления данных пользователя
func UpdateUser(user *models.User) error {
	result := initializers.DB.Save(&user)
	return result.Error
}

// DeleteUser Функция для удаления пользователя
func DeleteUser(userId uint) error {
	result := initializers.DB.Delete(&models.User{}, userId)
	return result.Error
}
