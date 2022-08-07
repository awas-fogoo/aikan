package routes

import (
	"awesomeProject0511/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	// 解决不同源
	r.Use(middleware.CORSMiddleware())
	v1 := r.Group("/api/v1")
	{
		HomeRouter(v1)
		ChannelRouter(v1)
	}

	return r
}
