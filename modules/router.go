package modules

import (
	"fmt"
	"net/http"

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
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})

	router := initRouter.Group("/api")
	{
		router.POST("signup")
	}

	return initRouter
}

// interceptor 와 exception 구현 필요
