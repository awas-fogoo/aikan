package controller

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"one/model"
	"testing"
)

func TestUploadVideo(t *testing.T) {
	dsn := "root:OwKO8HuAr0@tcp(74.48.75.173:3306)/one?charset=utf8mb4&parseTime=True&loc=Local"

	// 尝试连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connected successfully")
	}
	err = db.AutoMigrate(&model.VideoURL{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection handle: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	} else {
		log.Println("Database ping successful")
	}
	fmt.Println("----------------")
	// Create series data
	seriesDescription := "A thrilling sci-fi adventure."
	series := model.Series{
		Title:         "Space Explorers",
		Description:   &seriesDescription,
		Category:      new(string),
		TotalSeasons:  1,
		TotalEpisodes: 10,
	}
	result := db.Create(&series)
	if result.Error != nil {
		t.Errorf("Failed to create series: %v", result.Error)
	}
	// 应用同样的错误检查方法到其他的 db.Create 调用

	// Create season data
	seasonDescription := "The first and epic season of Space Explorers."
	season := model.Season{
		SeriesID:     series.ID,
		SeasonNumber: 1,
		Description:  &seasonDescription,
	}
	result2 := db.Create(&season)
	if result2.Error != nil {
		t.Errorf("Failed to create series: %v", result2.Error)
	}
	// 应用同样的错误检查方法到其他的 db.Create 调用

	// Create episodes and associated videos
	for i := 1; i <= 10; i++ {
		episodeTitle := fmt.Sprintf("Episode %d", i)
		episodeDesc := "What happens next will shock you."
		videoDesc := "High quality 1080p"
		coverUrl := "http://example.com/cover.jpg"
		video := model.Video{
			Title:           episodeTitle,
			Description:     &videoDesc,
			Uploader:        new(string),
			Duration:        45 * 60, // 45 minutes
			Category:        new(string),
			Resolution:      new(string),
			BelongsToSeries: &series.ID,
			CoverImageUrl:   &coverUrl,
			VideoURLs: []model.VideoURL{
				{URL: "http://example.com/video1.mp4"},
				{URL: "http://example.com/video2.mp4"},
			},
		}
		result3 := db.Create(&video)
		if result3.Error != nil {
			t.Errorf("Failed to create series: %v", result3.Error)
		}
		episode := model.Episode{
			SeasonID:      season.ID,
			EpisodeNumber: i,
			Title:         episodeTitle,
			Description:   &episodeDesc,
			Duration:      45 * 60, // 45 minutes
			VideoID:       video.ID,
		}
		result4 := db.Create(&episode)
		if result4.Error != nil {
			t.Errorf("Failed to create series: %v", result4.Error)
		}
	}

	// Create tags and associate them with videos
	tags := []string{"Sci-Fi", "Adventure", "Space"}
	var tagRecords []model.Tag
	for _, tagName := range tags {
		tag := model.Tag{TagName: tagName}
		db.FirstOrCreate(&tag, model.Tag{TagName: tagName})
		tagRecords = append(tagRecords, tag)
	}

	for _, tag := range tagRecords {
		db.Create(&model.VideoTag{
			VideoID: 1, // Assuming the video ID is 1 for demonstration
			TagID:   tag.ID,
		})
	}
}
