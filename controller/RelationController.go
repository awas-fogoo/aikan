package controller

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
)

func Relation(c *gin.Context) {
	uid := c.Query("uid")
	vid := c.Query("vid")
	//res := server.IsRedisLike(rdb, ctx, vid, uid)
	//cont := server.LikeIdCount(rdb, ctx, vid)

	db := common.InitDB()
	var videoRelationVo []vo.VideoRelationVo
	db.Model(&model.VideoRelation{}).Where("uid = ? and vid = ?", uid, vid).Scan(&videoRelationVo)

	//statusServer := server.StatusServer(db, vid, uid)
	c.JSON(200, dto.RetStruct{
		Ret:  true,
		Data: gin.H{"relation": videoRelationVo},
		Code: 0,
		Msg:  "",
	})
}
