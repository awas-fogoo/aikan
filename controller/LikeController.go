package controller

import (
	"awesomeProject0511/common"
	"awesomeProject0511/server"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
)

func Like(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, _ := common.ParseToken(tokenString)
	uid := claims.UserId
	videoVo := vo.VideoVo{}
	c.Bind(&videoVo)
	//vid := c.PostForm("vid")
	vid := videoVo.Vid
	res := server.LikeServer(vid, uid)
	c.JSON(200, res)

}
