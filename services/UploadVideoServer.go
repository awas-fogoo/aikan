package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/util"
	"bytes"
	"github.com/jinzhu/gorm"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (VideoService) UploadVideoServer(
	title, desc, url, coverUrl, tags string,
	categoryId, userId uint,
) (uint, *dto.RetDTO) {
	db := common.DB
	var user model.User

	if err := db.Where("id =?", userId).First(&user).Error; err != nil {
		return 0, dto.Error(4000, "用户不存在")
	}

	// 验证输入数据
	if len(title) == 0 {
		return 0, dto.Error(4000, "标题不能为空")
	}

	if len(desc) == 0 {
		return 0, dto.Error(4000, "描述不能为空")
	}

	if len(tags) == 0 {
		return 0, dto.Error(4000, "标签不能为空")
	}

	if categoryId <= 0 {
		return 0, dto.Error(4000, "无效的分类")
	}

	if !util.IsValidURL(url) {
		return 0, dto.Error(4000, "无效的视频 URL")
	}

	if !util.IsValidURL(coverUrl) {
		return 0, dto.Error(4000, "无效的封面 URL")
	}

	tagNames := handleTags(tags)
	// 获取所有分类
	var categories []model.Category
	err := db.Find(&categories).Error
	if err != nil {
		return 0, dto.Error(5000, "获取分类列表失败")
	}

	// 获取用户选择的分类
	var category model.Category
	err = db.Where("id = ?", categoryId).First(&category).Error
	if err != nil {
		return 0, dto.Error(5000, "获取用户选择的分类失败")
	}

	// 创建一个 Video 类型的变量
	video := model.Video{
		Title:       title,
		Description: desc,
		CategoryID:  category.ID,
		UserID:      userId,
	}

	// 创建新的视频标签关联
	if err := createVideoWithTags(db, &video, tagNames); err != nil {
		return 0, dto.Error(5000, "创建视频标签关联失败")
	}

	// 将文件上传到云端或者保存到本地，获取文件的 URL
	video.Url = url
	video.CoverUrl = coverUrl

	if strings.HasSuffix(url, ".mp4") {
		// 获取视频时长
		video.Duration = getVideoDuration(url)

		// 获取分辨率大小
		video.Quality = getVideoResolution(url)
	} else {
		video.Duration = 0
		video.Quality = "Unknown"
	}

	// 将视频关联到用户
	user.Videos = append(user.Videos, video)

	// 保存更改
	if err := db.Save(&user).Error; err != nil {
		return 0, dto.Error(5000, "保存用户视频失败")
	}
	return video.ID, nil
}

// CreateVideoWithTags 创建视频和标签，并建立多对多关系
func createVideoWithTags(db *gorm.DB, video *model.Video, tagNames []string) error {
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
