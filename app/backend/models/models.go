package models

import (
	"database/sql"
	"time"
)

type User struct {
	UserId           uint       `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username         string     `gorm:"size:20;not null" json:"username"`
	Email            string     `gorm:"size:60;unique;not null" json:"email"`
	PasswordHash     string     `gorm:"size:60;not null" json:"password_hash"`
	RegistrationDate time.Time  `gorm:"not null;autoCreateTime" json:"registration_date"`
	AccountStatus    string     `gorm:"size:255;not null;check:account_status IN ('active', 'inactive', 'pending');default:'active'" json:"account_status"`
	Inactive         *time.Time `gorm:"defalut:null" json:"inactive"`
	Gender           string     `gorm:"size:255;not null;check:gender IN ('male', 'female', 'prefer_not_to_say');default:prefer_not_to_say" json:"gender"`
	Biography        string     `gorm:"type:text;not null;default:''" json:"biography"`
	OnlineStatus     string     `gorm:"size:255;not null;check:online_status IN ('online', 'offline', 'away');default:'offline'" json:"online_status"`
	IsBanned         bool       `gorm:"not null; default:false" json:"is_banned"`
}

type Message struct {
	MessageID   uint          `gorm:"primaryKey;autoIncrement" json:"message_id"`
	SenderID    sql.NullInt32 `json:"sender_id"`
	RecipientID sql.NullInt32 `json:"recipient_id"`
	Message     string        `gorm:"type:text;not null" json:"message"`
	Timestamp   time.Time     `gorm:"not null" json:"timestamp"`
}

type MatchList struct {
	MatchID        uint `gorm:"primaryKey;autoIncrement" json:"chat_id"`
	FirstPersonID  uint `gorm:"not null;" json:"first_id"`
	SecondPersonID uint `gorm:"not null" json:"second_id"`
	User1          User `gorm:"foreignKey:FirstPersonID;constraint:onDelete:CASCADE" json:"-"`
	User2          User `gorm:"foreignKey:SecondPersonID;constraint:onDelete:CASCADE" json:"-"`
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
	// User             User      `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE;" json:"user"`
}

type Like struct {
	LikeID      uint `gorm:"primaryKey;autoIncrement" json:"like_id"`
	SenderID    uint `gorm:"not null;constraint:OnDelete:SET DEFAULT;" json:"sender_id"`
	RecipientID uint `gorm:"not null;constraint:OnDelete:SET DEFAULT;" json:"recipient_id"`
}
