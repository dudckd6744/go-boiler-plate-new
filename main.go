package main

import (
	"github.com/dudckd6744/go-boiler-plate/core"
	"github.com/dudckd6744/go-boiler-plate/modules"
)

func main() {

	core.DBConnection()
	router := modules.Router()

	router.Run(":5050")
}
