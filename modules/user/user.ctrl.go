package user

import (
	"github.com/dudckd6744/go-boiler-plate/common/response"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	result := Service.GetUser()
	response.Success(ctx, result)
}
