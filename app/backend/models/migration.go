package models

import "gorm.io/gorm"

func MigrateAll(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&User{},
		&Message{},
		&LabelInfo{},
		&UserLabel{},
		&BlockData{},
		&Block{},
		&Notification{},
		&Like{},
		&MatchList{},
	); err != nil {
		return err
	}
	return nil
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func MigrateMessage(db *gorm.DB) error {
	return db.AutoMigrate(&Message{})
}

func MigrateLabelInfo(db *gorm.DB) error {
	return db.AutoMigrate(&LabelInfo{})
}

func MigrateUserLabel(db *gorm.DB) error {
	return db.AutoMigrate(&UserLabel{})
}

func MigrateBlockData(db *gorm.DB) error {
	return db.AutoMigrate(&BlockData{})
}

func MigrateBlock(db *gorm.DB) error {
	return db.AutoMigrate(&Block{})
}

func MigrateNotification(db *gorm.DB) error {
	return db.AutoMigrate(&Notification{})
}

func MigrateLike(db *gorm.DB) error {
	return db.AutoMigrate(&Like{})
}

func MigrateMatchList(db *gorm.DB) error {
	return db.AutoMigrate(&MatchList{})
}
