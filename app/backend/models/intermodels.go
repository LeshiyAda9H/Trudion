package models

type UserPage struct {
	UserId       uint     `json:"user_id"`
	Username     string   `gorm:"size:20;not null" json:"username"`
	Gender       string   `gorm:"size:255;not null;check:gender IN ('male', 'female', 'prefer_not_to_say');default:prefer_not_to_say" json:"gender"`
	Biography    string   `gorm:"type:text;not null;default:' '" json:"biography"`
	Labels       []string `gorm:"type:text[]" json:"label"`
	OnlineStatus string   `gorm:"size:255;not null;check:online_status IN ('online', 'offline', 'away');default:'offline'" json:"online_status"`
}

type UserNotification struct {
	Message string
}
