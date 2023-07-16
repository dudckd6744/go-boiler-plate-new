package user

import (
	"github.com/dudckd6744/go-boiler-plate/common/response"
	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {

	response.Success(ctx, gin.H{"zz": "Gg"})
	return
}

func UserController(router gin.RouterGroup) {

	router.GET("/")

}
