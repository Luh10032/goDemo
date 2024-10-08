package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeUse() gin.HandlerFunc {

	return func(c *gin.Context) {
		t := time.Now()

		c.Next()
		cost := time.Since(t).Milliseconds()
		fmt.Printf("cost time:%d milliseconds \n", cost)
	}
}
