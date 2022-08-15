package controller

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/server"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, _ := common.ParseToken(tokenString)
	uid := claims.UserId
	videoVo := vo.VideoVo{}
	c.Bind(&videoVo)
	rdb := common.InitCache()
	ctx := common.Ctx
	//vid := c.PostForm("vid")
	vid := videoVo.Vid
	res := server.IsLike(rdb, ctx, vid, uid)
	cont := server.LikeIdCount(rdb, ctx, vid)
	c.JSON(200, dto.RetStruct{
		Ret: true,
		Data: gin.H{
			"like": res,
			"cont": cont,
		},
		Code: 0,
		Msg:  "",
	})
}
