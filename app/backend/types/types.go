package types

type VerifyEmailPayload struct {
	Email string `json:"email" binding:"required"`
}

type SignInPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpPayload struct {
	Username  string   `json:"username" binding:"required"`
	Email     string   `json:"email" binding:"required"`
	Password  string   `json:"password" binding:"required"`
	Gender    string   `json:"gender" binding:"required"`
	Label     []string `json:"label"`
	Biography string   `gorm:"default:''" json:"biography"`
}
