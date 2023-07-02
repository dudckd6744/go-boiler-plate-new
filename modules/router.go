package modules

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	initRouter := gin.Default()

	initRouter.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})

	router := initRouter.Group("/api")
	{
		router.POST("signup")
	}

	return initRouter
}
