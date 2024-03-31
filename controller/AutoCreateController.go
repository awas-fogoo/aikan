package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"one/common"
	"one/dto"
	"one/model"
	"one/util"
)

func AutoCreateUser(c *gin.Context) {

	addUser(c)
	addCategory(c)
	addTag(c)

}

func addUser(c *gin.Context) {
	db := common.DB
	var users []*model.User
	err := util.WithTransaction(db, func(tx *gorm.DB) error {
		// 增删改操作
		for i := 0; i < 20; i++ {
			user := &model.User{
				Username:  fmt.Sprintf("user%d", i),
				Nickname:  fmt.Sprintf("nickname%d", i),
				Email:     fmt.Sprintf("user%d@bup.pub", i),
				AvatarUrl: fmt.Sprintf("https://bup.pub/%d.jpg", i),
				Gender:    "男",
				Age:       uint(i),
			}
			user.Auth.Password = fmt.Sprintf("$2a$10$JfgQTe0bp.Xgd3zmYcHOwObCQ2nC6eMNIppT2z/jwdQonFZbs4rnK")
			if err := tx.Create(user).Error; err != nil {
				return err
			}
			users = append(users, user)
		}
		return nil
	})
	if err != nil {
		// 发生错误，需要回滚事务
		log.Fatalf(fmt.Sprintf("Failed to execute transactions: %v", err))
	}
	c.JSON(200, dto.Success(users))
	return
}

// 创建电影分类及其子分类
func addCategory(c *gin.Context) {
	db := common.DB
	err := util.WithTransaction(db, func(tx *gorm.DB) error {
		// init category
		// 建立分类
		categories := []model.Category{
			{Name: "番剧", Description: "以日本动画为代表的中国互联网用户所特别喜欢的分类、日本动画、国产动画、欧美动画等的长篇连载剧集"},
			{Name: "国创", Description: "以国产动画为代表的中国互联网用户所特别喜欢的分类"},
			{Name: "电影", Description: "各种类型的电影，如剧情片、喜剧片、动作片、恐怖片、科幻片等"},
			{Name: "电视剧", Description: "国内外的各种类型的电视剧，如古装剧、都市剧、悬疑剧、警匪剧等"},
			{Name: "综艺", Description: "真人秀、竞技类、脱口秀、娱乐综艺等"},
			{Name: "动画", Description: "日本动画、国产动画、欧美动画等"},
			{Name: "舞蹈", Description: "舞蹈教学、表演、比赛等"},
			{Name: "娱乐", Description: "各种综艺节目、娱乐活动、明星八卦等"},
			{Name: "美食", Description: "各种美食的制作方法、美食评测、美食分享等"},
			{Name: "时尚", Description: "时尚资讯、潮流服饰、美妆护肤等"},
			{Name: "旅游", Description: "各地景点、旅游攻略、旅游体验等"},
			{Name: "生活", Description: "日常生活中的各种实用技能、家居装修、宠物养护等"},
			{Name: "资讯", Description: "新闻资讯、时政评论、财经要闻、社会热点等"},
			{Name: "亲子", Description: "亲子教育、亲子互动、儿童教育等相关视频"},
			{Name: "知识", Description: "各种知识科普、学科知识、历史人文、学术讲座、读书笔记等等"},
			{Name: "影视", Description: "电影、电视剧、纪录片等影视作品的推荐、点评等"},
			{Name: "游戏", Description: "各种游戏攻略、游戏解说、游戏评测等相关视频"},
			{Name: "汽车", Description: "各种汽车新车试驾、汽车保养技巧等相关视频"},
			{Name: "财经", Description: "各种财经新闻、股票投资、理财规划等相关视频"},
			{Name: "萌宠", Description: "各种可爱的宠物、宠物养护、宠物日常等"},
			{Name: "运动", Description: "各种体育运动、健身训练、运动装备等"},
			{Name: "音乐", Description: "各种音乐类型的推荐、歌曲排行、音乐节目等"},
			{Name: "短片", Description: "各种类型的短片"},
			{Name: "科技", Description: "科技类视频，如科普、IT等相关内容"},
			{Name: "健康", Description: "健身、营养、医学、保健等"},
			{Name: "纪实类", Description: "纪录片、纪实类视频，包括自然、历史、探险、人物等内容"},
			{Name: "其他", Description: "其他类型的视频，无法归类到以上分类中的视频"},
		}
		// 为添加分类
		for i := range categories {
			err := tx.Create(&categories[i]).Error
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
	if err != nil {
		// 发生错误，需要回滚事务
		log.Fatalf(fmt.Sprintf("Failed to execute transactions: %v", err))
	}
	c.JSON(200, dto.Success("category data add  success"))
	return
}

func addTag(c *gin.Context) {
	db := common.DB
	err := util.WithTransaction(db, func(tx *gorm.DB) error {
		// 创建新的标签
		tags := []model.Tag{
			{Name: "动作"},
			{Name: "冒险"},
			{Name: "喜剧"},
			{Name: "犯罪"},
			{Name: "爱情"},
			{Name: "科幻"},
			{Name: "悬疑"},
			{Name: "恐怖"},
			{Name: "灾难"},
			{Name: "神秘"},
			{Name: "励志"},
			{Name: "纪录片"},
			{Name: "奇幻"},
			{Name: "战争"},
			{Name: "童话"},
			{Name: "音乐"},
			{Name: "运动"},
			{Name: "西部"},
			{Name: "青春"},
			{Name: "武侠"},
			{Name: "轻喜剧"},
			{Name: "历史"},
			{Name: "玄幻"},
			{Name: "竞技"},
			{Name: "古装"},
			{Name: "家庭"},
			{Name: "谍战"},
			{Name: "社会"},
			{Name: "都市"},
			{Name: "科幻动作"},
			{Name: "青春校园"},
		}
		// 为添加标签
		for i := range tags {
			err := tx.Create(&tags[i]).Error
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
	if err != nil {
		// 发生错误，需要回滚事务
		log.Fatalf(fmt.Sprintf("Failed to execute transactions: %v", err))
	}
	c.JSON(200, dto.Success("category data add  success"))
	return
}
