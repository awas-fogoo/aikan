package middleware

import (
	"awesomeProject0511/common"
	"awesomeProject0511/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get Authorization header
		tokenString := c.GetHeader("Authorization")

		//validate token formate 判断是否以bearer开头 Abort返回不向下执行函数了
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 402, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 从7开始是因为上Bearer 是7位
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 402, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 验证通过获取claims中的user.id,并且查询这个id的信息
		userId := claims.UserId
		DB := common.InitDB()
		var user model.User
		DB.First(&user, userId)

		//用户是否存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 402, "msg": "权限不足"})
			c.Abort()
			return
		}

		//存在之后将user的信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
