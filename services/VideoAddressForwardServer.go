package services

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

func requestVideo(videoURL string, start int64, end int64) (io.ReadCloser, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", videoURL, nil)
	if err != nil {
		return nil, err
	}
	rangeHeader := fmt.Sprintf("bytes=%d-%d", start, end)
	req.Header.Set("Range", rangeHeader)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusPartialContent {
		resp.Body.Close()
		return nil, fmt.Errorf("server does not support partial content")
	}

	return resp.Body, nil
}

func VideoStreamServer(c *gin.Context) {
	// 获取要处理的视频资源地址，并进行 Base64 解码
	url := c.Query("url")
	decodedURL, err := base64.StdEncoding.DecodeString(url)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid URL parameter")
		return
	}
	// 解析 Range 头部字段
	rangeHeader := c.Request.Header.Get("Range")
	start, end := parseRange(rangeHeader)
	// 发起视频请求
	resp, err := requestVideo(string(decodedURL), start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 返回视频响应内容
	c.Status(http.StatusPartialContent)
	c.Header("Content-Type", "video/mp4")
	c.Header("Accept-Ranges", "bytes")
	c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/*", start, end))
	io.Copy(c.Writer, resp)
	resp.Close()
}

func parseRange(rangeHeader string) (start int64, end int64) {
	r := regexp.MustCompile(`bytes=(\d*)-(\d*)`)
	matches := r.FindStringSubmatch(rangeHeader)
	if len(matches) != 3 {
		return 0, 0
	}
	startStr, endStr := matches[1], matches[2]
	if startStr == "" {
		end, _ = strconv.ParseInt(endStr, 10, 64)
		start = -1 * end
	} else if endStr == "" {
		start, _ = strconv.ParseInt(startStr, 10, 64)
	} else {
		start, _ = strconv.ParseInt(startStr, 10, 64)
		end, _ = strconv.ParseInt(endStr, 10, 64)
	}
	return start, end
}
