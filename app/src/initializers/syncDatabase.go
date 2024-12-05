package initializers

import "src/models"

func SyncDatabase() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Message{},
		&models.LabelInfo{},
		&models.UserLabel{},
		&models.BlockData{},
		&models.Block{},
		&models.Notification{},
		&models.Like{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
}
