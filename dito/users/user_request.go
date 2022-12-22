package usersdito

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Email    string `gorm:"type : varchar(255)" json:"email" validate:"required,email_exist"`
	Password string `gorm : "type : varchar(255)" json:"password" validate:"required"`
	Address  string `gorm :"type : varchar (255)" json:"address" validate:"required"`
	FullName string `gorm: "type: varchar(255)" json:"name" validate:"required"`
	Phone    string `gorm : "type: varchar(255)" json:"phone" validate:"required"`
}
