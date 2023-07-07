package modules

import (
	"fmt"

	"github.com/dudckd6744/go-boiler-plate/common/response"
	middleware "github.com/dudckd6744/go-boiler-plate/core/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Exception occurred:", r)

		}
	}()

	initRouter := gin.Default()

	initRouter.Use(middleware.LoggingMiddleware())

	initRouter.GET("ping", func(ctx *gin.Context) {

		response.Success(ctx, gin.H{"success": true})
	})

	router := initRouter.Group("/api")
	{
		router.POST("signup")
	}

	return initRouter
}

// interceptor 와 exception 구현 필요
