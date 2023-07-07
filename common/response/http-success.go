package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data gin.H) {

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"error":  nil,
			"data":   data})
}
