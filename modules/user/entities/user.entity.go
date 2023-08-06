package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"type:varchar(25)" json:"name"`
	Email     string         `gorm:"type:varchar(50)" json:"email"`
	Password  string         `gorm:"type:varchar(100)" json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
