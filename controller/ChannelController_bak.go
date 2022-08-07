package controller

//
//import (
//	"awesomeProject0511/common"
//	"awesomeProject0511/model"
//	"awesomeProject0511/vo"
//	"github.com/gin-gonic/gin"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//	"net/http"
//	"strconv"
//)
//
//type DataAll struct {
//	ContentList []vo.ChannelListVo `json:"contentList"`
//	UpRecommend []vo.UpRecommendVo `json:"upRecommend"`
//}
//
//func GetTopChannelDetail(c *gin.Context) {
//	videoUrl := c.Param("id")
//	db := common.InitDB()
//	defer db.Close()
//
//	var videoInfo model.ChannelList
//	var recInfo []model.ChannelList
//	db.First(&videoInfo, "video_url=?", videoUrl)
//	var list []string
//	var listRec []string
//
//	createAtTime := videoInfo.CreatedAt.Format("2006-01-02 15:04:05")
//	list = append(
//		list,
//		createAtTime,
//		videoInfo.ImgUrl,
//		videoInfo.Hot,
//		videoInfo.Title,
//		videoInfo.HeadUrl,
//		videoInfo.Name,
//		strconv.Itoa(int(videoInfo.Fans)),
//		strconv.FormatBool(videoInfo.Follow),
//		strconv.Itoa(int(videoInfo.See)),
//		strconv.Itoa(int(videoInfo.Danmu)),
//		videoInfo.VideoUrl,
//		videoInfo.Desc,
//		strconv.Itoa(int(videoInfo.Like)),
//		strconv.Itoa(int(videoInfo.Dislike)),
//		strconv.Itoa(int(videoInfo.Collection)),
//		strconv.Itoa(int(videoInfo.Share)),
//		strconv.Itoa(int(videoInfo.Comment)),
//		videoInfo.Uid,
//	)
//	//fmt.Println(list)
//
//	db.Find(&recInfo)
//
//	rows := make([]map[string]interface{}, len(recInfo))
//	for i, finger := range recInfo {
//		rows[i] = map[string]interface{}{
//			"Uid":      finger.Uid,
//			"ImgUrl":   finger.ImgUrl,
//			"VideoUrl": finger.VideoUrl,
//			"Title":    finger.Title,
//		}
//	}
//	for _, row := range rows {
//		listRec = append(listRec, row["Uid"].(string), row["ImgUrl"].(string), row["VideoUrl"].(string), row["Title"].(string))
//	}
//
//	var y = DataAll{
//		[]vo.ChannelListVo{
//			{list[0], list[1], list[2], list[3], list[4], list[5], list[6],
//				list[7], list[8], list[9], list[10], list[11], list[12],
//				list[13], list[14], list[15], list[16], list[17]},
//		},
//		[]vo.UpRecommendVo{
//			{listRec[8], listRec[9], listRec[10], listRec[11]},
//			{listRec[12], listRec[13], listRec[14], listRec[15]},
//			{listRec[0], listRec[1], listRec[2], listRec[3]},
//			{listRec[4], listRec[5], listRec[6], listRec[7]},
//			{listRec[16], listRec[17], listRec[18], listRec[19]},
//		}}
//	c.JSON(http.StatusOK, gin.H{
//		"ret": true,
//		"data": gin.H{
//			"contentList": y.ContentList,
//			"upRecommend": y.UpRecommend,
//		},
//	})
//}

//type ChannelList struct {
//	gorm.Model
//	Cover       string  `gorm:"varchar(255);not null" json:"imgUrl"`
//	Hot         string  `gorm:"varchar(20)" json:"hot"`
//	Title       string  `gorm:"varchar(100);not null" json:"title"`
//	Name        string  `gorm:"varchar(100);not null" json:"name"`
//	Videos      string  `gorm:"varchar(255);not null;unique" json:"videoUrl"`
//	Uid         string  `gorm:"varchar(100);not null" json:"uid"`
//	Copyright   bool    `gorm:"not null"`  //是否为原创(版权)
//	Weights     float32 `gorm:"default:0"` //视频权重(目前还没使用)
//	Clicks      int     `gorm:"default:0"` //点击量
//	Review      int     `gorm:"not null"`  //审核状态
//	PartitionID uint    `gorm:"default:0"` //分区ID
//
//	HeadUrl    string `gorm:"varchar(255);not null" json:"headUrl"`
//	Desc       string `gorm:"varchar(255);not null" json:"desc"`
//	Fans       uint   `gorm:"int(100)" json:"fans"`
//	Follow     bool   `gorm:"default:0" json:"follow"`
//	See        uint   `gorm:"int(255);not null;default:1" json:"see"`
//	Danmu      uint   `gorm:"int(200)" json:"danmu"`
//	Like       uint   `gorm:"int(255)" json:"like"`
//	Dislike    uint   `gorm:"int(100)" json:"dislike"`
//	Collection uint   `gorm:"int(200)" json:"collection"`
//	Share      uint   `gorm:"int(50);" json:"share"`
//	Comment    uint   `gorm:"int(255)" json:"comment"`
//}
