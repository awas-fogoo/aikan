package controller

import (
	"awesomeProject0511/common"
	"awesomeProject0511/server"
	"awesomeProject0511/vo"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Like(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, _ := common.ParseToken(tokenString)
	fmt.Println(claims)
	videoVo := vo.VideoVo{}
	c.Bind(&videoVo)
	//vid := c.PostForm("vid")
	//vid := videoVo.Vid
	uid := videoVo.Uid
	vid := videoVo.Vid
	res := server.LikeServer(uid, vid)
	c.JSON(200, res)

}
