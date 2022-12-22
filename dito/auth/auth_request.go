package authDito

type LoginRequest struct {
	Email    string `gorm:"type : varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type : varchar(255)" json:"password" validate:"required"`
}
