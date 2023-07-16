package modules

import (
	"fmt"

	"github.com/dudckd6744/go-boiler-plate/common/response"
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

	router := initRouter.Group("/api")

	{
		router.Group("/user", User)
	}
	// 유저 구조체를 불러들여서 json 으로 반환되는지 확인 필요

	return initRouter
}
