package models

import (
	"fmt"
	"sesi_6/pkg/helpers"
	"time"

	// valid "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	FullName  string         `json:"full_name" valid:"required-Your full name is required"`
	Username  string         `json:"username" valid:"required-Your username is required"`
	Email     string         `json:"email" valid:"required-Your email is required,email-Your email is not valid"`
	Password  *string        `json:"password" valid:"required-Your password is required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type RegisterRequest struct {
	FullName string `json:"full_name" valid:"required"`
	Username string `json:"username" valid:"required"`
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
}

type LoginRequest struct {
	Username string `json:"username" valid:"required-Your username is required"`
	Password string `json:"password" valid:"required-Your password is required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        User   `json:"user"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Printf("user %+v", u)

	// _, err = valid.ValidateStruct(u)

	// if err != nil {
	// 	return err
	// }

	getPassword := *u.Password

	hashedPassword, err := helpers.HashPassword(getPassword)
	if err != nil {
		return err
	}

	u.Password = &hashedPassword

	return nil
}
