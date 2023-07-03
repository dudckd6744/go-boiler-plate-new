package interceptors

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 인터셉트 로직
		start := time.Now()

		// 다음 핸들러 호출
		c.Next()

		end := time.Now()

		fmt.Println(end.Sub(start).Seconds(), "초")
	}
}
