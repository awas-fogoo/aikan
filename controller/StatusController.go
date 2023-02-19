package controller

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
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

	//vid := c.PostForm("vid")
	vid := videoVo.Vid
	//res := server.IsRedisLike(rdb, ctx, vid, uid)
	//cont := server.LikeIdCount(rdb, ctx, vid)

	db := common.InitDB()
	var count int64
	db.Model(&model.ChannelLiked{}).Where("vid = ? and status = ?", vid, 1).Count(&count)
	// SELECT count(1) FROM users WHERE name = 'jinzhu'; (count)

	statusServer := server.StatusServer(db, vid, uid)
	c.JSON(200, dto.RetStruct{
		Ret: true,
		Data: gin.H{
			"cont":         count,
			"statusServer": statusServer,
		},
		Code: 0,
		Msg:  "",
	})
}
