package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"fmt"
)

func (VideoService) AddDanmu(vid, uid uint, start, duration uint64, colour, prior bool, content, mode string, style model.DanmukuStyle) error {
	// 连接数据库
	db := common.InitDB()
	defer db.Close()
	fmt.Println(style)
	danmu := model.Danmuku{
		VideoID:  vid,
		Content:  content,
		Start:    start,
		Duration: duration,
		Prior:    prior,
		Colour:   colour,
		Mode:     mode,
		Style:    style,
		UserID:   uid,
	}

	// 插入弹幕记录
	if err := db.Create(&danmu).Error; err != nil {
		return err
	}

	return nil
}

type DanmukuRaw struct {
	ID              uint   `json:"id"`
	UserID          uint   `json:"user_id"`
	Content         string `json:"txt"`
	Start           uint64 `json:"start"`
	Duration        uint64 `json:"duration"`
	Prior           bool   `json:"prior"`
	Colour          bool   `json:"colour"`
	Mode            string `json:"mode"`
	Color           string `json:"color"`
	FontSize        string `json:"font_size"`
	Border          string `json:"border"`
	BorderRadius    string `json:"border_radius"`
	Padding         string `json:"padding"`
	BackgroundColor string `json:"background_color"`
}

func (VideoService) GetDanmu(vid string) ([]vo.DanmukuResponseVo, error) {
	// 连接数据库
	db := common.InitDB()
	defer db.Close()

	// 执行原始查询
	rows, err := db.Raw("SELECT id, user_id, start, duration, prior, colour, content, mode, color, font_size, border, border_radius, padding, background_color FROM danmukus WHERE video_id = ?", vid).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var danmuRes []vo.DanmukuResponseVo

	// 扫描结果并映射到 DanmukuResponseVo 结构体
	for rows.Next() {
		var raw DanmukuRaw
		if err := db.ScanRows(rows, &raw); err != nil {
			return nil, err
		}

		danmu := vo.DanmukuResponseVo{
			UserID:   raw.UserID,
			Content:  raw.Content,
			Start:    raw.Start,
			Duration: raw.Duration,
			Prior:    raw.Prior,
			Colour:   raw.Colour,
			Mode:     raw.Mode,
			Style: struct {
				Color           string `json:"color"`
				FontSize        string `json:"fontSize"`
				Border          string `json:"border"`
				BorderRadius    string `json:"borderRadius"`
				Padding         string `json:"padding"`
				BackgroundColor string `json:"backgroundColor"`
			}{
				Color:           raw.Color,
				FontSize:        raw.FontSize,
				Border:          raw.Border,
				BorderRadius:    raw.BorderRadius,
				Padding:         raw.Padding,
				BackgroundColor: raw.BackgroundColor,
			},
		}

		danmuRes = append(danmuRes, danmu)
	}

	return danmuRes, err
}
