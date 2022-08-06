package controller

import (
	"awesomeProject0511/common"
	"awesomeProject0511/model"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Data struct {
	SwiperList    []model.SwiperList `json:"swiperList"`
	PopRecommend  string             `json:"popRecommend"`
	ClassicReview string             `json:"classicReview"`
}

func GetHomeList(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()

	var data model.SwiperList

	var users []string
	res, _ := db.Raw("select id,uid,img_url,video_href from swiper_lists").Rows()
	for res.Next() {
		res.Scan(&data.Id, &data.VideoHref, &data.Uid, &data.ImgUrl)
		// fmt.Printf("data: %v\n", data)
		users = append(users, data.Id, data.VideoHref, data.Uid, "channel/"+data.ImgUrl)
	}
	y := Data{[]model.SwiperList{
		{users[0], users[1], users[2], users[3]},
		{users[4], users[5], users[6], users[7]},
		{users[8], users[9], users[10], users[11]},
		{users[12], users[13], users[14], users[15]}},
		"热门推荐", "大家常看"}

	//f.SwiperList = append(f.SwiperList, y)
	c.JSON(200, gin.H{
		"ret":  true,
		"data": y,
	})

}
