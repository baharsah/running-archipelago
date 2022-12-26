package authdito

type LoginRequest struct {
	Email    string `gorm:"type : varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type : varchar(255)" json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email    string `gorm:"type : varchar(255)" json:"email" validate:"required,email_exist"`
	Password string `gorm : "type : varchar(255)" json:"password" validate:"required"`
	Address  string `gorm :"type : varchar (255)" json:"address" validate:"required"`
	FullName string `gorm: "type: varchar(255)" json:"name" validate:"required"`
	Phone    string `gorm : "type: varchar(255)" json:"phone" validate:"required"`
}
