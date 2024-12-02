package models

import (
	"time"
)

type User struct {
	UserId           uint       `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username         string     `gorm:"size:20;not null" json:"username"`
	Email            string     `gorm:"size:60;unique;not null" json:"email"`
	PasswordHash     string     `gorm:"size:60;not null" json:"password_hash"`
	RegistrationDate time.Time  `gorm:"not null;autoCreateTime" json:"registration_date"`
	AccountStatus    string     `gorm:"size:255;not null;check:account_status IN ('active', 'inactive', 'pending')" json:"account_status" default:"active"`
	Inactive         *time.Time `gorm:"" json:"inactive" default:"nil"`
	Gender           string     `gorm:"size:255;not null;check:gender IN ('male', 'female', 'prefer_not_to_say')" json:"gender" default:"prefer_not_to_say"`
	Biography        string     `gorm:"type:text;not null" json:"biography" default:" "`
	OnlineStatus     string     `gorm:"size:255;not null;check:online_status IN ('online', 'offline', 'away')" json:"online_status" default:"offline"`
	IsBanned         bool       `gorm:"not null" json:"is_banned" default:"false"`
}

type Message struct {
	MessageID  uint      `gorm:"primaryKey;autoIncrement" json:"message_id"`
	SenderID   uint      `gorm:"not null" json:"sender_id"`
	ReceiverID uint      `gorm:"not null" json:"receiver_id"`
	Message    string    `gorm:"type:text;not null" json:"message"`
	Timestamp  time.Time `gorm:"not null" json:"timestamp"`
}

type LabelInfo struct {
	LabelID   uint   `gorm:"primaryKey;autoIncrement" json:"label_id"`
	LabelName string `gorm:"size:30;not null" json:"label_name"`
}

type UserLabel struct {
	UserID  uint      `gorm:"not null" json:"user_id"`
	LabelID uint      `gorm:"not null" json:"label_id"`
	User    User      `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
	Label   LabelInfo `gorm:"foreignKey:LabelID;constraint:onDelete:CASCADE" json:"label"`
}

type BlockData struct {
	UserID      uint      `gorm:"primaryKey;not null" json:"user_id"`
	BannedDate  time.Time `gorm:"not null" json:"banned_date"`
	Description string    `gorm:"type:text;not null" json:"description" default:"Violating platform rules"`
	User        User      `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
}

type Block struct {
	BlockID       uint `gorm:"primaryKey;autoIncrement" json:"block_id"`
	BlockerID     uint `gorm:"not null;constraint:OnDelete:SET DEFAULT;" json:"blocker_id"`
	BlockedUserID uint `gorm:"not null;constraint:OnDelete:SET DEFAULT;" json:"blocked_user_id"`
}

type Notification struct {
	NotificationID   uint      `gorm:"primaryKey;autoIncrement" json:"notification_id"`
	UserID           uint      `gorm:"not null" json:"user_id"`
	Message          string    `gorm:"type:text;not null" json:"message"`
	IsRead           bool      `gorm:"not null" json:"is_read" default:"false"`
	NotificationDate time.Time `gorm:"not null;autoCreateTime" json:"notification_date"`
	User             User      `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE;" json:"user"`
}

type Like struct {
	LikeID        uint   `gorm:"primaryKey;autoIncrement" json:"like_id"`
	FirstLikerID  uint   `gorm:"not null;constraint:OnDelete:SET DEFAULT;" json:"first_liker_id"`
	SecondLikerID uint   `gorm:"not null;constraint:OnDelete:SET DEFAULT;" json:"second_liker_id"`
	LikeStatus    string `gorm:"type:varchar(255);not null" json:"like_status"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
