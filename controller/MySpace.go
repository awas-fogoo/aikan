package controller

import (
	"awesomeProject0511/common"
	"awesomeProject0511/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataMySpace struct {
	MySpace []model.MySpace `json:"mySpace"`
}

func Myspace(c *gin.Context) {
	uid := c.Param("id")
	db := common.InitDB()
	defer db.Close()
	var data model.User
	var dataCon model.ContentList
	var list []string
	db.First(&data, "uid=?", uid)
	db.First(&dataCon, "uid=?", uid)
	//db.Where("uid = ?", uid).First(&model.MySpace{})
	list = append(list, data.Name, data.Uid, dataCon.ImgUrl, dataCon.VideoUrl, dataCon.Title)
	v := DataMySpace{[]model.MySpace{{list[0], list[1], list[2], list[3], list[4]}}}
	c.JSON(http.StatusOK, gin.H{
		"ret": true,
		"data": gin.H{
			"mySpace": v.MySpace,
		},
	})
}
