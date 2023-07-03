package main

import (
	"github.com/dudckd6744/go-boiler-plate/core"
	"github.com/dudckd6744/go-boiler-plate/modules"
)

func main() {

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Exception occurred:", r)

	// 	}
	// }()

	core.DBConnection()
	router := modules.Router()

	router.Run(":5050")
}
