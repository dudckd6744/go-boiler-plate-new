package user

import (
	"time"

	"github.com/dudckd6744/go-boiler-plate/core"
	"gorm.io/gorm"
)

type UserRepository struct{}

var Repository UserRepository

// User struct에 따라 select가 설정된다.
// json = alias를 의미한다.
type User struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (r *UserRepository) GetUser() *[]User {

	user := []User{}

	core.DBConnection().Table("user").Find(&user)

	return &user
}
