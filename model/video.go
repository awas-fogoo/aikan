package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*
	mongodb

用户行为数据：
维持一个UserBehavior的集合，包含每个用户的喜欢、不喜欢、收藏、观看历史、搜索历史以及关注与被关注的数据。
这样的数据结构有助于个性化服务和安全性，因为所有与个人用户相关的行为都封装在单独的文档中。

视频统计数据：
维持一个VideoStatistics的集合，专门用于跟踪和管理每个视频的观看次数、喜欢数、不喜欢数和收藏数。
这种独立的统计表便于执行高效的更新和查询操作，特别是在内容浏览和推荐算法中。
*/
type UserBehavior struct {
	UserID        string          `bson:"user_id"`
	Followings    []Following     `bson:"followings"`     //关注
	Followers     []Follower      `bson:"followers"`      //粉丝
	Likes         []Like          `bson:"likes"`          //喜欢
	Dislikes      []Dislike       `bson:"dislikes"`       //不喜欢
	Favorites     []Favorite      `bson:"favorites"`      //收藏
	WatchHistory  []WatchHistory  `bson:"watch_history"`  //观看记录
	SearchHistory []SearchHistory `bson:"search_history"` //搜索记录
}
type Following struct {
	FollowedID string    `bson:"followed_id"`
	Timestamp  time.Time `bson:"timestamp"`
}

type Follower struct {
	FollowerID string    `bson:"follower_id"`
	Timestamp  time.Time `bson:"timestamp"`
}
type Like struct {
	VideoID   string    `bson:"video_id"`
	Timestamp time.Time `bson:"timestamp"`
}

type Dislike struct {
	VideoID   string    `bson:"video_id"`
	Timestamp time.Time `bson:"timestamp"`
}

type Favorite struct {
	VideoID   string    `bson:"video_id"`
	Timestamp time.Time `bson:"timestamp"`
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

// VideoStatistics 视频信息聚合
type VideoStatistics struct {
	VideoID       string    `bson:"video_id"`
	WatchCount    int64     `bson:"watch_count"`
	LikeCount     int64     `bson:"like_count"`
	DislikeCount  int64     `bson:"dislike_count"`
	FavoriteCount int64     `bson:"favorite_count"`
	Timestamp     time.Time `bson:"timestamp"`
}

// VideoAnalytics 视频详情
type VideoAnalytics struct {
	VideoID       string   `bson:"video_id"`
	KeyFrames     []string `bson:"key_frames"`
	Subtitles     []string `bson:"subtitles"`
	IsSeries      bool     `bson:"is_series"`
	SeasonNumber  *int     `bson:"season_number,omitempty"`  // 使用指针允许字段为空
	EpisodeNumber *int     `bson:"episode_number,omitempty"` // 使用指针允许字段为空
}

// UserFeedbacks 用户反馈
type UserFeedbacks struct {
	VideoID  string    `bson:"video_id"`
	UserID   string    `bson:"user_id"`
	Rating   *float64  `bson:"rating"`
	Comments []Comment `bson:"comments"`
}

// SeriesAnalytics 系列分析
type SeriesAnalytics struct {
	SeriesID       string    `bson:"series_id"`
	TotalWatchTime int       `bson:"total_watch_time"`
	AverageRating  float64   `bson:"average_rating"`
	Comments       []Comment `bson:"comments"`
}

// Comment 评论
type Comment struct {
	CommentID primitive.ObjectID `bson:"comment_id"`
	Content   string             `bson:"content"`
	Timestamp time.Time          `bson:"timestamp"`
}

// ContextData 分享
type ContextData struct {
	UserID             string               `bson:"user_id"`
	VideoID            string               `bson:"video_id"`
	WatchDatetime      time.Time            `bson:"watch_datetime"`
	SocialInteractions []SocialInteractions `bson:"interactions"` // 分享到哪个平台的行为
}

type SocialInteractions struct {
	ActionType     string    `bson:"action_type"`     // type:share
	TargetPlatform string    `bson:"target_platform"` // QQ/WeChat/Twitter
	Timestamp      time.Time `bson:"timestamp"`
}

// Danmuku 弹幕库
type Danmuku struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    uint               `bson:"user_id"`
	Content   string             `bson:"txt"`
	Start     uint64             `bson:"start"`
	Duration  uint64             `bson:"duration"`
	Prior     bool               `bson:"prior"`
	Colour    bool               `bson:"color"`
	Mode      string             `bson:"mode"`
	Style     DanmukuStyle       `bson:"style"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

// DanmukuStyle 弹幕库样式
type DanmukuStyle struct {
	Color           string `bson:"color"`
	FontSize        string `bson:"fontSize"`
	Border          string `bson:"border"`
	BorderRadius    string `bson:"borderRadius"`
	Padding         string `bson:"padding"`
	BackgroundColor string `bson:"backgroundColor"`
}
