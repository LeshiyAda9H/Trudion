package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"src/config"
	"src/models"
)

var DB *gorm.DB

// InitializeDatabase Инициализация соединения с базой данных
func InitializeDatabase() (*gorm.DB, error) {
	config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DbConfig.Host,
		config.DbConfig.User,
		config.DbConfig.Password,
		config.DbConfig.DBName,
		config.DbConfig.Port,
		config.DbConfig.SSLMode,
		config.DbConfig.TimeZone,
	)

	var err error
	// Инициализируем соединение с базой данных
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}
	return DB, nil
}

// CreateUser Функция для создания нового пользователя
func CreateUser(user *models.User) error {
	// Создаем нового пользователя в базе данных
	result := DB.Create(&user)
	return result.Error
}

// FindUserByEmail Функция для поиска пользователя по email
func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser Функция для обновления данных пользователя
func UpdateUser(user *models.User) error {
	result := DB.Save(&user)
	return result.Error
}

// DeleteUser Функция для удаления пользователя
func DeleteUser(userId uint) error {
	result := DB.Delete(&models.User{}, userId)
	return result.Error
}
