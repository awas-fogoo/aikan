package controller

import (
	"awesomeProject0511/common"
	"awesomeProject0511/model"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

type DataAll struct {
	ContentList []model.ContentListTo `json:"contentList"`
	UpRecommend []model.UpRecommendTo `json:"upRecommend"`
}

func GetTopChannelDetail(c *gin.Context) {
	videoUrl := c.Param("id")
	db := common.InitDB()
	defer db.Close()

	var videoInfo model.ContentList
	var recInfo []model.ContentList
	db.First(&videoInfo, "video_url=?", videoUrl)
	var list []string
	var listRec []string

	createAtTime := videoInfo.CreatedAt.Format("2006-01-02 15:04:05")
	list = append(
		list,
		createAtTime,
		videoInfo.ImgUrl,
		videoInfo.Hot,
		videoInfo.Title,
		videoInfo.HeadUrl,
		videoInfo.Name,
		strconv.Itoa(int(videoInfo.Fans)),
		strconv.FormatBool(videoInfo.Follow),
		strconv.Itoa(int(videoInfo.See)),
		strconv.Itoa(int(videoInfo.Danmu)),
		videoInfo.VideoUrl,
		videoInfo.Desc,
		strconv.Itoa(int(videoInfo.Like)),
		strconv.Itoa(int(videoInfo.Dislike)),
		strconv.Itoa(int(videoInfo.Collection)),
		strconv.Itoa(int(videoInfo.Share)),
		strconv.Itoa(int(videoInfo.Comment)),
		videoInfo.Uid,
	)
	//fmt.Println(list)

	db.Find(&recInfo)

	rows := make([]map[string]interface{}, len(recInfo))
	for i, finger := range recInfo {
		rows[i] = map[string]interface{}{
			"Uid":      finger.Uid,
			"ImgUrl":   finger.ImgUrl,
			"VideoUrl": finger.VideoUrl,
			"Title":    finger.Title,
		}
	}
	for _, row := range rows {
		listRec = append(listRec, row["Uid"].(string), row["ImgUrl"].(string), row["VideoUrl"].(string), row["Title"].(string))
	}

	var y = DataAll{
		[]model.ContentListTo{
			{list[0], list[1], list[2], list[3], list[4], list[5], list[6],
				list[7], list[8], list[9], list[10], list[11], list[12],
				list[13], list[14], list[15], list[16], list[17]},
		},
		[]model.UpRecommendTo{
			{listRec[8], listRec[9], listRec[10], listRec[11]},
			{listRec[12], listRec[13], listRec[14], listRec[15]},
			{listRec[0], listRec[1], listRec[2], listRec[3]},
			{listRec[4], listRec[5], listRec[6], listRec[7]},
			{listRec[16], listRec[17], listRec[18], listRec[19]},
		}}
	c.JSON(http.StatusOK, gin.H{
		"ret": true,
		"data": gin.H{
			"contentList": y.ContentList,
			"upRecommend": y.UpRecommend,
		},
	})
}
