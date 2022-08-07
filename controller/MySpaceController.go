package controller

import (
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
)

type DataMySpace struct {
	MySpace []vo.MySpaceVo `json:"mySpace"`
}

func Myspace(c *gin.Context) {
	//uid := c.Param("id")
	//db := common.InitDB()
	//defer db.Close()
	//var data model.User
	//var dataCon model.ChannelList
	//var list []string
	//db.First(&data, "uid=?", uid)
	//db.First(&dataCon, "uid=?", uid)
	////db.Where("uid = ?", uid).First(&model.MySpace{})
	//list = append(list, data.Name, data.Uid, dataCon.ImgUrl, dataCon.VideoUrl, dataCon.Title)
	//v := DataMySpace{[]vo.MySpaceVo{{list[0], list[1], list[2], list[3], list[4]}}}
	//c.JSON(http.StatusOK, gin.H{
	//	"ret": true,
	//	"data": gin.H{
	//		"mySpace": v.MySpace,
	//	},
	//})
}
