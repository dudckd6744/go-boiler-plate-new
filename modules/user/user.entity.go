package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(25)"`
	Email    string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(100)"`
}
