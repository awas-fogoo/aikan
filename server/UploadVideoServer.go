package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/util"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func UploadVideoServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	getUser, _ := c.Get("user")
	userDto := dto.ToUserDTO(getUser.(model.User))
	var user model.User
	db.Where("id =?", userDto.ID).First(&user)
	title := c.PostForm("title")
	desc := c.PostForm("description")
	url := c.PostForm("url")
	coverUrl := c.PostForm("cover_url")
	tags := c.PostForm("tags")
	categoryId := util.StringToUint(c.PostForm("category_id"))

	if len(title) == 0 && len(desc) == 0 && categoryId <= 0 && util.IsValidURL(url) && util.IsValidURL(url) {
		c.JSON(0, dto.Error(-1, "cannot be empty or url acquisition failed"))
		return
	}
	tagNmaes := handleTags(tags)
	// 获取所有分类
	var categories []model.Category
	err := db.Find(&categories).Error
	if err != nil {
		log.Fatalln(err)
		return
	}
	// 获取用户选择的分类
	var category model.Category
	err = db.Where("id = ?", categoryId).First(&category).Error
	if err != nil {
		log.Println(err)
		c.JSON(0, dto.Error(0, "category does not exist"))
		return
	}

	// 创建一个 Video 类型的变量
	video := model.Video{
		Title:       title,
		Description: desc,
		CategoryID:  category.ID,
		UserID:      userDto.ID,
	}
	// 创建新的视频标签关联
	if err := CreateVideoWithTags(db, &video, tagNmaes); err != nil {
		log.Fatal(err)
	}
	// 将文件上传到云端或者保存到本地，获取文件的 URL
	video.Url = url
	video.CoverUrl = coverUrl

	// 获取视频时长
	video.Duration = getVideoDuration(url)

	// 获取分辨率大小
	video.Quality = getVideoResolution(url)
	// 将视频关联到用户
	user.Videos = append(user.Videos, video)

	// 保存更改
	err = db.Save(&user).Error
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(0, dto.Success(&video))
}

// CreateVideoWithTags 创建视频和标签，并建立多对多关系
func CreateVideoWithTags(db *gorm.DB, video *model.Video, tagNames []string) error {
	// 创建视频
	if err := db.Create(video).Error; err != nil {
		return err
	}

	// 查找或创建标签
	var tags []model.Tag
	for _, name := range tagNames {
		var tag model.Tag
		db.Where("name = ?", name).FirstOrCreate(&tag, model.Tag{Name: name})
		tags = append(tags, tag)
	}

	// 建立视频和标签的多对多关系
	for _, tag := range tags {
		db.Create(&model.VideoTag{VideoID: video.ID, TagID: tag.ID})
	}

	return nil
}

func handleTags(tags string) []string {
	separators := []string{",", ";", "，", "；", "\"", "”", "“", "'", "‘", "、", "/"}
	for _, sep := range separators {
		tags = strings.ReplaceAll(tags, sep, ",")
	}
	tagList := strings.Split(tags, ",")
	var result []string
	for _, tag := range tagList {
		tag = strings.TrimSpace(tag)
		if tag != "" && !contains(result, tag) {
			result = append(result, tag)
		}
	}
	return result
}

func contains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

func getVideoResolution(url string) string {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=s=x:p=0", url)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Println(err)
		return "Unknown"
	}
	re := regexp.MustCompile(`(\d+)x(\d+)`)
	match := re.FindStringSubmatch(out.String())
	if len(match) == 3 {
		width, _ := strconv.Atoi(match[1])
		if width >= 3840 {
			return "4K"
		} else if width >= 1920 {
			return "1080p"
		} else if width >= 1280 {
			return "720p"
		} else if width >= 600 {
			return "360p"
		} else if width >= 480 {
			return "360p"
		} else {
			return "Unknown"
		}
	} else {
		return "Unknown"
	}
}

func getVideoDuration(url string) (duration float64) {
	maxRetries := 3                  // 最大重试次数
	retryInterval := 1 * time.Second // 重试间隔
	defaultDuration := 0             // 默认时长

	cmd := exec.Command("ffprobe", "-i", url, "-show_entries", "format=duration", "-v", "quiet", "-of", "csv=p=0")

	var output []byte
	var err error

	for i := 0; i < maxRetries; i++ {
		output, err = cmd.CombinedOutput()
		if err == nil {
			break // 如果没有错误，退出循环
		}

		log.Printf("Command failed with error: %v. Retrying in %v...", err, retryInterval)
		time.Sleep(retryInterval) // 等待一段时间后重试
	}

	if err != nil {
		log.Printf("Command failed with error: %v", err)
		return 0
	}

	// 解析视频时长
	durationString := string(output)
	durationString = strings.Trim(durationString, "\n")
	durationFloat, err := strconv.ParseFloat(durationString, 64)
	if err != nil {
		log.Printf("Failed to parse video duration: %v", err)
		duration = float64(defaultDuration)
	} else {
		duration = durationFloat
	}

	if duration <= 0 {
		log.Printf("Invalid video duration: %v. Using default duration of %v instead.", duration, defaultDuration)
		duration = float64(defaultDuration)
	}

	return duration
}
