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
	dsn := "root:OwKO8HuAr0@tcp(74.48.75.174:3306)/one?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connected successfully")
	}

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

	// 创建标签
	tag1 := model.Tag{TagName: "动作"}
	tag2 := model.Tag{TagName: "古装"}
	db.Create(&tag1)
	db.Create(&tag2)

	// 创建分类
	category1 := model.Category{Name: "剧情", Order: 1}
	category2 := model.Category{Name: "冒险", Order: 2}
	db.Create(&category1)
	db.Create(&category2)

	// 使用事务进行操作
	err = db.Transaction(func(tx *gorm.DB) error {
		// 创建系列数据
		actor := "张若昀-范闲,李沁-林婉儿,陈道明-庆帝"
		year := 2024
		detailDescription := "故事讲述了一个身世神秘的青年范闲，历经家族、江湖、庙堂的种种考验与锤炼，他秉持正义、良善，开始了新的人生征途，继续书写出这段不同寻常又酣畅淋漓的人生传奇。 剧作既根植于传统文化，又超脱于传统历史小说，是一部极具东方古典气韵和现代意识的力作，致力弘扬珍惜当下美好，不忘初心的中华传统价值美德。余年有幸，与君再相逢。"
		var detail = model.Detail{
			Title:          "庆余年第二季",
			Description:    &detailDescription,
			Categories:     1,
			CoverImageUrl:  nil,
			Director:       nil,
			Scriptwriter:   nil,
			CurrentEpisode: 1,
			TotalEpisodes:  35,
			Actors:         &actor,
			RegionID:       1,
			Year:           &year,
			Tags:           []model.Tag{tag1, tag2},
		}
		if err := tx.Create(&detail).Error; err != nil {
			return err
		}

		// 创建集数和关联视频
		for i := 1; i <= 8; i++ {
			detailTitle := fmt.Sprintf("Episode %d", i)
			videoDesc := "High quality 1080p"
			coverUrl := "http://example.com/cover.jpg"

			video := model.Video{
				DetailID:      detail.ID,
				Title:         detailTitle,
				Description:   &videoDesc,
				CoverImageUrl: &coverUrl,
				Tags:          []model.Tag{tag1, tag2},
				VideoURLs: []model.VideoURL{
					{
						Order:    1,
						URL:      "http://example.com/video1.mp4",
						Source:   nil,
						Quality:  nil,
						Language: nil,
						Subtitle: nil,
					},
					{
						Order:    2,
						URL:      "http://example.com/video2.mp4",
						Source:   nil,
						Quality:  nil,
						Language: nil,
						Subtitle: nil,
					},
				},
			}
			if err := tx.Create(&video).Error; err != nil {
				return err
			}
		}

		// 如果所有操作都成功，提交事务
		return nil
	})

	// 如果事务中任何操作失败，事务将被回滚
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	} else {
		log.Println("Transaction completed successfully")
	}
}
