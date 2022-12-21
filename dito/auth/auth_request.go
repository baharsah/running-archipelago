package auth

type LoginRequest struct {
	Email    string `gorm:"type : varchar(255)" json:"name" validate:"required"`
	Password string `gorm:"type : varchar(255)" json:"password" validate:"required"`
}
