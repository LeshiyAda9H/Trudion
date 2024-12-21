package initializers

import (
	"errors"
	"fmt"
	"src/models"

	"gorm.io/gorm"
)

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
		&models.MatchList{},
	)
	if err != nil {
		panic("failed to migrate database")
	}

	// fill label_info with default values
	insertDefaultLabels()
}

func insertDefaultLabels() {
	labels := []models.LabelInfo{
		{LabelName: "football"},
		{LabelName: "basketball"},
		{LabelName: "volleyball"},
		{LabelName: "painting"},
		{LabelName: "music"},
		{LabelName: "photography"},
		{LabelName: "programming"},
		{LabelName: "gaming"},
		{LabelName: "robotics"},
		{LabelName: "physics"},
		{LabelName: "chemistry"},
		{LabelName: "biology"},
		{LabelName: "movies"},
		{LabelName: "books"},
		{LabelName: "travel"},
	}

	for _, label := range labels {
		var existingLabel models.LabelInfo
		if err := DB.Where("label_name = ?", label.LabelName).First(&existingLabel).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := DB.Create(&label).Error; err != nil {
					fmt.Printf("Failed to create label %s: %v\n", label.LabelName, err)
				}
			} else {
				fmt.Printf("Error while checking label %s: %v\n", label.LabelName, err)
			}
		}
	}
}
