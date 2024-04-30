package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*  mongodb  */

type UserBehavior struct {
	UserID          string          `bson:"user_id"`
	WatchHistory    []WatchHistory  `bson:"watch_history"`
	SearchHistory   []SearchHistory `bson:"search_history"`
	InteractionData []Interaction   `bson:"interaction_data"`
}

type WatchHistory struct {
	VideoID   string    `bson:"video_id"`
	WatchTime time.Time `bson:"watch_time"`
	Duration  int       `bson:"duration"`
}

type SearchHistory struct {
	SearchTerm string    `bson:"search_term"`
	SearchTime time.Time `bson:"search_time"`
}

type Interaction struct {
	VideoID    string    `bson:"video_id"`
	ActionType string    `bson:"action_type"`
	Timestamp  time.Time `bson:"timestamp"`
}

type VideoAnalytics struct {
	VideoID       string   `bson:"video_id"`
	KeyFrames     []string `bson:"key_frames"`
	Subtitles     []string `bson:"subtitles"`
	IsSeries      bool     `bson:"is_series"`
	SeasonNumber  *int     `bson:"season_number,omitempty"`  // 使用指针允许字段为空
	EpisodeNumber *int     `bson:"episode_number,omitempty"` // 使用指针允许字段为空
}
type UserFeedbacks struct {
	VideoID  string    `bson:"video_id"`
	UserID   string    `bson:"user_id"`
	Rating   float64   `bson:"rating"`
	Comments []Comment `bson:"comments"`
}

type SeriesAnalytics struct {
	SeriesID       string    `bson:"series_id"`
	TotalWatchTime int       `bson:"total_watch_time"`
	AverageRating  float64   `bson:"average_rating"`
	Comments       []Comment `bson:"comments"`
}

type Comment struct {
	Content   string    `bson:"content"`
	Timestamp time.Time `bson:"timestamp"`
}

type ContextData struct {
	UserID             string               `bson:"user_id"`
	VideoID            string               `bson:"video_id"`
	WatchDatetime      time.Time            `bson:"watch_datetime"`
	SocialInteractions []SocialInteractions `bson:"interactions"`
}

type SocialInteractions struct {
	ActionType     string    `bson:"action_type"`
	TargetPlatform string    `bson:"target_platform"`
	Timestamp      time.Time `bson:"timestamp"`
}

type Danmuku struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   uint               `bson:"user_id"`
	Content  string             `bson:"txt"`
	Start    uint64             `bson:"start"`
	Duration uint64             `bson:"duration"`
	Prior    bool               `bson:"prior"`
	Colour   bool               `bson:"color"`
	Mode     string             `bson:"mode"`
	Style    DanmukuStyle       `bson:"style"`
}

type DanmukuStyle struct {
	Color           string `bson:"color"`
	FontSize        string `bson:"fontSize"`
	Border          string `bson:"border"`
	BorderRadius    string `bson:"borderRadius"`
	Padding         string `bson:"padding"`
	BackgroundColor string `bson:"backgroundColor"`
}

// Story 结构体与 a_story 表对应
type Story struct {
	ID         uint   `gorm:"primaryKey" json:"id"`                  // id 主键
	VideoType  string `gorm:"column:video_type" json:"video_type"`   // 视频类型
	Name       string `gorm:"column:name" json:"name"`               // 剧情
	CreateBy   string `gorm:"column:create_by" json:"create_by"`     // 创建人
	CreateTime string `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateBy   string `gorm:"column:update_by" json:"update_by"`     // 修改人
	UpdateTime string `gorm:"column:update_time" json:"update_time"` // 修改时间
	Remark     string `gorm:"column:remark" json:"remark"`           // 备注
}
