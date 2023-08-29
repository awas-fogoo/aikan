package services

import (
	"fmt"
	"testing"
)

func TestUploadVideo(t *testing.T) {
	// 设置测试的输入数据
	title := "Test Video"
	desc := "Test Description"
	url := "https://example.com/test.mp4"
	coverUrl := "https://example.com/test.jpg"
	tags := "tag1, tag2"
	categoryId := uint(1)
	userId := uint(123)

	// 调用服务层函数
	videoURL, retDTO := UploadVideo(title, desc, url, coverUrl, tags, categoryId, userId)
	// 测试无效的 URL
	//invalidURL := "invalid_url"
	// 测试标题为空
	//emptyTitle := ""
	fmt.Println(videoURL)
	fmt.Println(retDTO)
}
