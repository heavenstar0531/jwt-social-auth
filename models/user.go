package models

import (
	"jwt-go/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Your password must be at least 6 characters"`
	Email    string    `gorm:"not null;unique" json:"email" form:"email" valid:"required~Your email is required, email~Your email is invalid" validate:"required,email"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	err = nil
	return
}
