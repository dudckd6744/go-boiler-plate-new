package main

import (
	"fmt"

	"github.com/dudckd6744/go-boiler-plate/core"
	"github.com/dudckd6744/go-boiler-plate/modules"
	"github.com/dudckd6744/go-boiler-plate/modules/user/entities"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Exception occurred:", r)

		}
	}()

	db := core.DBConnection()

	db.AutoMigrate(&entities.User{})

	router := modules.Router()

	router.Run(":5050")
}
