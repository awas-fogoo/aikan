package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"one/common"
	"one/dto"
	"one/model"
	"one/util"
)

func GetProfileServer(c *gin.Context) {
	db := common.DB
	tokenString := c.GetHeader("Authorization")
	viewer := uint(0) // 默认为未登录用户
	if tokenString != "" {
		// 从7开始是因为Bearer是7位
		tokenString = tokenString[7:]
		_, claims, err := common.ParseToken(tokenString)
		if err == nil {
			viewer = claims.UserId
		}
	}

	viewee := util.StringToUint(c.Query("uid"))

	// 访问者没有登入，并且没有输入uid
	if viewee <= 0 && viewer == 0 {
		c.JSON(200, dto.Error(4000, "Uid Get failed"))
		return
	}
	// 访问者登入了，但是没有输入要查看的uid，则是查看自己
	if viewee <= 0 && viewer != 0 {
		viewee = viewer
	}
	user := &model.User{}
	err := db.Preload("Videos").Preload("Following").Preload("Followers").First(user, viewee).Error
	if err != nil {
		log.Println(err)
		return
	}
	// 获取用户上传视频总数
	count := len(user.Videos)

	// 获取用户关注数量
	followingCount := len(user.Following)

	// 获取用户粉丝数量
	followerCount := len(user.Followers)

	// 查询登入的当前用户是否已关注目标用户
	isFollowing := false
	if viewer != 0 {
		for _, followee := range user.Following {
			if followee.ID == viewer {
				isFollowing = true
				break
			}
		}
	}

	// 按照最新上传的时间排序
	var videos []model.Video
	if err = db.Where("user_id = ?", viewee).Order("views DESC").Limit(12).Find(&videos).Error; err != nil {
		log.Println(err)
		return
	}
	// 将每个视频的 views 和 cover_url 添加到数组中
	results := make([]map[string]interface{}, 0, len(videos))
	for _, video := range videos {
		result := make(map[string]interface{})
		result["id"] = video.ID
		result["views"] = video.Views
		result["cover_url"] = video.CoverUrl
		results = append(results, result)
	}

	data := map[string]interface{}{
		"posts":          count,
		"followingCount": followingCount,
		"followerCount":  followerCount,
		"isFollowing":    isFollowing,
		"videos":         results,
		"aboutMe":        user.AboutMe,
		"avatarUrl":      user.AvatarUrl,
		"nickname":       user.Nickname,
		"username":       user.Username,
		"backgroundUrl":  user.BackgroundUrl,
	}

	c.JSON(200, dto.Success(data))

}
