package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomError(c *gin.Context, data gin.H) {

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusBadRequest,
			"error":  data,
			"data":   nil})
}
