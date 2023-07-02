package core

import (
	"github.com/dudckd6744/go-boiler-plate/modules/user"
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func DBConnection() *gorm.DB {

	dsn := "root:12341234@tcp(localhost:3306)/go_test"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Info),
	})

	if err == nil {
		color.Red("db connection success!!")
	}

	db.AutoMigrate(&user.User{})

	sqlDB, _ := db.DB()
	// 현재 사용되지 않는 연결을 유지
	sqlDB.SetMaxIdleConns(10)
	// 동시 연결수
	sqlDB.SetMaxOpenConns(50)

	return db
}
