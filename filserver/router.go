package filserver

import (
	"demo/filserver/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.Use(middleware.TimeUse())
	router(r)

	r.Run()
}

func router(r *gin.Engine) {
	r.POST("/create", Create)
	r.POST("/read", Read)

}
