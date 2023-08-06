package modules

import (
	"fmt"

	"github.com/dudckd6744/go-boiler-plate/common/response"
	"github.com/dudckd6744/go-boiler-plate/modules/user"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Exception occurred:", r)

		}
	}()

	initRouter := gin.Default()

	initRouter.GET("ping", func(ctx *gin.Context) {
		response.Success(ctx, gin.H{"success": true})
	})

	userRouter := initRouter.Group("api/user")

	{
		userRouter.GET("/", func(ctx *gin.Context) { user.GetUser(ctx) })

	}

	return initRouter
}
