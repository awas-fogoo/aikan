package web_interface

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"github.com/gin-gonic/gin"
)

func Like(c *gin.Context) {
	// 获取token里的所有信息
	user, _ := c.Get("user")

	// 类型断言 可以说成是类型强制转换
	uid := dto.ToUserDto(user.(model.User)).Uid
	vid := c.PostForm("vid")

	// 创建存储信息的数据库，没有则创建
	db := common.InitDB()
	defer db.Close()
	db.AutoMigrate(&model.Archive{})

	// 到这里应该是添加成功，添加成功之后，---需要把视频里的点赞数量加一，进行修改，先读取，加一，再修改
	// 读取数据库 这里读取的可能有点慢，后续需要修改
	var dataLike model.ContentList
	var archive model.Archive
	db.First(&archive, "like=?", vid)

	if archive.Like != vid {
		db.Model(&archive).Where("uid=?", uid).Update("like", "")
		db.First(&dataLike, "video_url=?", vid)
		numLike := dataLike.Like
		numLike -= 1
		db.Model(&dataLike).Where("video_url=?", vid).Update("like", numLike)
		c.JSON(200, gin.H{
			"msg":  "错误",
			"like": 0,
		})
		return
	}

	// 添加数据，数据拼接进去数据库，后面再修改，先添加一个
	var info = model.Archive{
		Follow:   "",
		History:  "",
		Like:     vid,
		Dislike:  "",
		Favorite: "",
		Time:     model.Time{},
		Uid:      uid,
	}
	db.Create(&info)

	db.First(&dataLike, "video_url=?", vid)
	numLike := dataLike.Like
	numLike += 1
	db.Model(&dataLike).Where("video_url=?", vid).Update("like", numLike)

	//
	c.JSON(200, gin.H{
		"msg":     "ok",
		"like":    1,
		"likenum": numLike,
	})
}
