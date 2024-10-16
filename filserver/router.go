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
	v1 := r.Group("/file")
	{
		v1.POST("/", Create)
		v1.GET("/:name", Read)
		v1.DELETE("/:name", Delete)
	}

}
