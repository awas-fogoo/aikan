package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	// 解决不同源
	//r.Use(middleware.CORSMiddleware())
	r.Use(cors.Default())

	v1 := r.Group("/api/v1")
	{
		AutoCreateUserRoute(v1)

		// 用户认证api
		UserAuthenticationRoute(v1)

		// 视频模块
		VideoRoute(v1)

		// 用户模块
		UserRoute(v1)
	}
	return r
}
