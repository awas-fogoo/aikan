package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"one/model"
	"testing"
	"time"
)

// GetFollowingCount - 获取用户关注的人数
func GetFollowingCount(ctx context.Context, collection *mongo.Collection, userID string) (int, error) {
	filter := bson.M{"user_id": userID}
	projection := bson.M{"_id": 0, "following_count": bson.M{"$size": "$followings"}}
	options := options.FindOne().SetProjection(projection)

	var result struct {
		FollowingCount int `bson:"following_count"`
	}

	err := collection.FindOne(ctx, filter, options).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("No user found with the given ID.")
			return 0, nil // 返回0表示没有错误，但用户不存在
		}
		log.Printf("Error retrieving following count: %v\n", err)
		return 0, err
	}

	return result.FollowingCount, nil
}

// GetFollowerCount - 获取关注该用户的人数
func GetFollowerCount(ctx context.Context, collection *mongo.Collection, userID string) (int, error) {
	filter := bson.M{"user_id": userID}
	projection := bson.M{"_id": 0, "follower_count": bson.M{"$size": "$followers"}}
	options := options.FindOne().SetProjection(projection)

	var result struct {
		FollowerCount int `bson:"follower_count"`
	}

	err := collection.FindOne(ctx, filter, options).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("No user found with the given ID.")
			return 0, nil // 返回0表示没有错误，但用户不存在
		}
		log.Printf("Error retrieving follower count: %v\n", err)
		return 0, err
	}

	return result.FollowerCount, nil
}

// GetLikes - 获取用户的喜欢列表
func GetLikes(ctx context.Context, collection *mongo.Collection, userID string) ([]model.Like, error) {
	filter := bson.M{"user_id": userID}
	var result model.UserBehavior

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("No documents found for the user.")
			return nil, nil // No documents found, return nil slice and no error.
		}
		log.Printf("Error retrieving likes: %v\n", err)
		return nil, err
	}

	return result.Likes, nil
}

// AddLike - 添加喜欢，并提供操作反馈
func AddLike(ctx context.Context, collection *mongo.Collection, userID, videoID string) error {
	like := model.Like{
		VideoID:   videoID,
		Timestamp: time.Now(),
	}
	update := bson.M{"$push": bson.M{"likes": like}}
	filter := bson.M{"user_id": userID}

	result, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 {
		log.Println("Like successfully added.")
	} else if result.UpsertedCount > 0 {
		log.Println("Like added and new document was created.")
	} else {
		log.Println("No like added or document updated.")
	}

	return nil
}

// RemoveLike - 移除喜欢，并检查操作结果
func RemoveLike(ctx context.Context, collection *mongo.Collection, userID, videoID string) error {
	filter := bson.M{"user_id": userID, "likes.video_id": videoID}
	update := bson.M{"$pull": bson.M{"likes": bson.M{"video_id": videoID}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		log.Println("No documents matched the filter, nothing to update.")
	} else if result.ModifiedCount == 0 {
		log.Println("Document found but no modifications made (element might not exist).")
	} else {
		log.Println("Like successfully removed.")
	}

	return nil
}

// GetDislikes - 获取用户的不喜欢列表
func GetDislikes(ctx context.Context, collection *mongo.Collection, userID string) ([]model.Dislike, error) {
	filter := bson.M{"user_id": userID}
	var result model.UserBehavior

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("No documents found for the user.")
			return nil, nil // No documents found, return nil slice and no error.
		}
		log.Printf("Error retrieving dislikes: %v\n", err)
		return nil, err
	}

	return result.Dislikes, nil
}

// AddDislike - 添加不喜欢
func AddDislike(ctx context.Context, collection *mongo.Collection, userID, videoID string) error {
	dislike := model.Dislike{
		VideoID:   videoID,
		Timestamp: time.Now(),
	}
	update := bson.M{"$push": bson.M{"dislikes": dislike}}
	filter := bson.M{"user_id": userID}

	result, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 {
		log.Println("Dislike successfully added.")
	} else if result.UpsertedCount > 0 {
		log.Println("Dislike added and new document was created.")
	} else {
		log.Println("No dislike added or document updated.")
	}

	return nil
}

// RemoveDislike - 移除不喜欢
func RemoveDislike(ctx context.Context, collection *mongo.Collection, userID, videoID string) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{"$pull": bson.M{"dislikes": bson.M{"video_id": videoID}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		log.Println("No dislike found or modification made.")
	} else {
		log.Println("Dislike successfully removed.")
	}

	return nil
}

// GetFollowings - 获取用户关注的人列表
func GetFollowings(ctx context.Context, collection *mongo.Collection, userID string) ([]model.Following, error) {
	filter := bson.M{"user_id": userID}
	var userBehavior model.UserBehavior

	err := collection.FindOne(ctx, filter).Decode(&userBehavior)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("No user found with the given ID.")
			return nil, nil // 返回nil表示没有错误，但也没有找到文档
		}
		log.Printf("Error retrieving followings: %v\n", err)
		return nil, err
	}

	return userBehavior.Followings, nil
}

// AddFollowing - 添加关注，并提供操作反馈
func AddFollowing(ctx context.Context, collection *mongo.Collection, userID, followedID string) error {
	following := model.Following{
		FollowedID: followedID,
		Timestamp:  time.Now(),
	}
	update := bson.M{"$push": bson.M{"followings": following}}
	filter := bson.M{"user_id": userID}

	result, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 {
		log.Println("Following successfully added.")
	} else if result.UpsertedCount > 0 {
		log.Println("Following added and new document was created.")
	} else {
		log.Println("No following added or document updated.")
	}

	return nil
}

// RemoveFollowing - 移除关注，并检查操作结果
func RemoveFollowing(ctx context.Context, collection *mongo.Collection, userID, followedID string) error {
	filter := bson.M{"user_id": userID, "followings.followed_id": followedID}
	update := bson.M{"$pull": bson.M{"followings": bson.M{"followed_id": followedID}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		log.Println("No documents matched the filter, nothing to update.")
	} else if result.ModifiedCount == 0 {
		log.Println("Document found but no modifications made (element might not exist).")
	} else {
		log.Println("Following successfully removed.")
	}

	return nil
}

// GetFollowers - 获取关注该用户的人列表
func GetFollowers(ctx context.Context, collection *mongo.Collection, userID string) ([]model.Follower, error) {
	filter := bson.M{"user_id": userID}
	var userBehavior model.UserBehavior

	err := collection.FindOne(ctx, filter).Decode(&userBehavior)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("No user found with the given ID.")
			return nil, nil // 返回nil表示没有错误，但也没有找到文档
		}
		log.Printf("Error retrieving followers: %v\n", err)
		return nil, err
	}

	return userBehavior.Followers, nil
}

// GetFavorites - 获取用户的收藏列表
func GetFavorites(ctx context.Context, collection *mongo.Collection, userID string) ([]model.Favorite, error) {
	filter := bson.M{"user_id": userID}
	var result model.UserBehavior

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("No documents found for the user.")
			return nil, nil // No documents found, return nil slice and no error.
		}
		log.Printf("Error retrieving favorites: %v\n", err)
		return nil, err
	}

	return result.Favorites, nil
}

// AddFavorite - 添加收藏，并提供操作反馈
func AddFavorite(ctx context.Context, collection *mongo.Collection, userID, videoID string) error {
	favorite := model.Favorite{
		VideoID:   videoID,
		Timestamp: time.Now(),
	}
	update := bson.M{"$push": bson.M{"favorites": favorite}}
	filter := bson.M{"user_id": userID}

	result, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 {
		log.Println("Favorite successfully added.")
	} else if result.UpsertedCount > 0 {
		log.Println("Favorite added and new document was created.")
	} else {
		log.Println("No favorite added or document updated.")
	}

	return nil
}

// RemoveFavorite - 移除收藏，并检查操作结果
func RemoveFavorite(ctx context.Context, collection *mongo.Collection, userID, videoID string) error {
	filter := bson.M{"user_id": userID, "favorites.video_id": videoID}
	update := bson.M{"$pull": bson.M{"favorites": bson.M{"video_id": videoID}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		log.Println("No documents matched the filter, nothing to update.")
	} else if result.ModifiedCount == 0 {
		log.Println("Document found but no modifications made (element might not exist).")
	} else {
		log.Println("Favorite successfully removed.")
	}

	return nil
}

// GetWatchHistory - 获取用户的观看历史
func GetWatchHistory(ctx context.Context, collection *mongo.Collection, userID string) ([]model.WatchHistory, error) {
	filter := bson.M{"user_id": userID}
	var result model.UserBehavior

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.WatchHistory, nil
}

// AddWatchHistory - 添加观看历史，并提供操作反馈
func AddWatchHistory(ctx context.Context, collection *mongo.Collection, userID, videoID string, duration int) error {
	watchHistory := model.WatchHistory{
		VideoID:   videoID,
		WatchTime: time.Now(),
		Duration:  duration,
	}
	update := bson.M{"$push": bson.M{"watch_history": watchHistory}}
	filter := bson.M{"user_id": userID}

	result, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 {
		log.Println("Watch history successfully added.")
	} else if result.UpsertedCount > 0 {
		log.Println("Watch history added and new document was created.")
	} else {
		log.Println("No watch history added or document updated.")
	}

	return nil
}

// RemoveWatchHistory - 移除观看历史
func RemoveWatchHistory(ctx context.Context, collection *mongo.Collection, userID, videoID string, watchTime time.Time) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{"$pull": bson.M{"watch_history": bson.M{"video_id": videoID, "watch_time": watchTime}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		log.Println("No watch history found or modification made.")
	} else {
		log.Println("Watch history successfully removed.")
	}

	return nil
}

// GetSearchHistory - 获取用户的搜索历史
func GetSearchHistory(ctx context.Context, collection *mongo.Collection, userID string) ([]model.SearchHistory, error) {
	filter := bson.M{"user_id": userID}
	var result model.UserBehavior

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.SearchHistory, nil
}

// AddSearchHistory - 添加搜索历史，并提供操作反馈
func AddSearchHistory(ctx context.Context, collection *mongo.Collection, userID, searchTerm string) error {
	searchHistory := model.SearchHistory{
		SearchTerm: searchTerm,
		SearchTime: time.Now(),
	}
	update := bson.M{"$push": bson.M{"search_history": searchHistory}}
	filter := bson.M{"user_id": userID}

	result, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 {
		log.Println("Search history successfully added.")
	} else if result.UpsertedCount > 0 {
		log.Println("Search history added and new document was created.")
	} else {
		log.Println("No search history added or document updated.")
	}

	return nil
}

// RemoveSearchHistory - 移除搜索历史
func RemoveSearchHistory(ctx context.Context, collection *mongo.Collection, userID, searchTerm string, searchTime time.Time) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{"$pull": bson.M{"search_history": bson.M{"search_term": searchTerm, "search_time": searchTime}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		log.Println("No search history found or modification made.")
	} else {
		log.Println("Search history successfully removed.")
	}

	return nil
}

func TestUserBehaviors(t *testing.T) {
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:yO08ZnSHHq@116.one:27017"))
	if err != nil {
		log.Fatal("连接到MongoDB失败：", err)
	}
	defer client.Disconnect(context.TODO())

	// 获取MongoDB集合的引用
	userBehaviorCollection := client.Database("One").Collection("user_behaviors")
	//GetLikes()
	//GetDislikes()
	//GetFavorites()
	//GetFollowerCount()
	//GetFollowers()
	//GetFollowings()
	//GetWatchHistory()
	//GetSearchHistory()
	//
	//RemoveWatchHistory()
	//RemoveDislike()
	//RemoveFavorite()
	//RemoveLike()
	//RemoveSearchHistory()
	//RemoveFollowing()

	videoAnalyticsCollection := client.Database("One").Collection("video_analytics")
	seriesAnalyticsCollection := client.Database("One").Collection("series_analytics")
	contextDataCollection := client.Database("One").Collection("context_data")
	userFeedbacksCollection := client.Database("One").Collection("user_feedbacks")
	danmukuCollection := client.Database("One").Collection("danmuku")
	//var result model.UserBehavior
	//filter := bson.D{{"user_id", "admin"}}
	//err = userBehaviorCollection.FindOne(context.TODO(), filter).Decode(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("找到的用户行为：%+v\n", result)

	// select
	//var results []model.VideoAnalytics
	//cursor, err := videoAnalyticsCollection.Find(context.TODO(), bson.D{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err = cursor.All(context.TODO(), &results); err != nil {
	//	log.Fatal(err)
	//}
	//for _, v := range results {
	//	fmt.Printf("找到的视频分析：%+v\n", v)
	//}

	// 伪造UserBehavior文档
	userBehavior := model.UserBehavior{
		UserID: "admin",
		WatchHistory: []model.WatchHistory{
			{
				VideoID:   "33",
				WatchTime: time.Now().Add(-48 * time.Hour), // 48小时前
				Duration:  120,
			},
		},
		SearchHistory: []model.SearchHistory{
			{
				SearchTerm: "funny cats",
				SearchTime: time.Now().Add(-24 * time.Hour), // 24小时前
			},
		},
		Likes: []model.Like{
			{
				VideoID:   "33",
				Timestamp: time.Now().Add(-24 * time.Hour),
			},
		},
		Dislikes: []model.Dislike{
			{
				VideoID:   "33",
				Timestamp: time.Now().Add(-24 * time.Hour),
			},
		},
		Favorites: []model.Favorite{
			{
				VideoID:   "33",
				Timestamp: time.Now().Add(-24 * time.Hour),
			},
		},
		Followings: []model.Following{
			{
				FollowedID: "user2",
				Timestamp:  time.Now().Add(-24 * time.Hour),
			},
		},
		Followers: []model.Follower{
			{
				FollowerID: "user3",
				Timestamp:  time.Now().Add(-24 * time.Hour),
			},
		},
	}

	// 伪造VideoAnalytics文档
	videoAnalytics := model.VideoAnalytics{
		VideoID:       "33",
		KeyFrames:     []string{"frame1.jpg", "frame2.jpg"},
		Subtitles:     []string{"Welcome to the video", "Thanks for watching"},
		IsSeries:      false,
		SeasonNumber:  nil,
		EpisodeNumber: nil,
	}

	// 伪造SeriesAnalytics文档
	seriesAnalytics := model.SeriesAnalytics{
		SeriesID:       "1",
		TotalWatchTime: 3600,
		AverageRating:  4.5,
		Comments: []model.Comment{
			{
				Content:   "Amazing series!",
				Timestamp: time.Now().Add(-72 * time.Hour),
			},
		},
	}

	// 伪造ContextData文档
	contextData := model.ContextData{
		UserID:        "admin",
		VideoID:       "33",
		WatchDatetime: time.Now().Add(-1 * time.Hour),
		SocialInteractions: []model.SocialInteractions{
			{
				ActionType:     "share",
				TargetPlatform: "Twitter",
				Timestamp:      time.Now().Add(-1 * time.Hour),
			},
		},
	}

	// Current timestamp
	currentTime := time.Now()

	// 伪造Danmuku数据
	danmukuData := model.Danmuku{
		UserID:   12345,
		Content:  "Hello, world!",
		Start:    100,
		Duration: 5000,
		Prior:    true,
		Colour:   true,
		Mode:     "scroll",
		Style: model.DanmukuStyle{
			Color:           "#FFFFFF",
			FontSize:        "14px",
			Border:          "1px solid #000",
			BorderRadius:    "4px",
			Padding:         "5px",
			BackgroundColor: "#000000",
		},
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	{
		// 插入UserFeedbacks文档到MongoDB
		userFeedbacks := model.UserFeedbacks{
			VideoID: "33",
			UserID:  "admin",
			Rating:  nil,
			Comments: []model.Comment{
				{
					CommentID: primitive.NewObjectID(),
					Content:   "Loved the video!",
					Timestamp: time.Now(),
				},
				{
					Content:   "Very informative.",
					Timestamp: time.Now(),
				},
			},
		}
		// 新评论
		newComment := model.Comment{
			Content:   "Another great aspect!",
			Timestamp: time.Now(),
		}

		// 更新操作
		update := bson.M{
			"$set": bson.M{
				"rating": userFeedbacks.Rating, // 可以添加更多需要更新的字段
			},
			"$push": bson.M{
				"comments": bson.M{"$each": []model.Comment{newComment}},
			},
		}

		// 更新过滤器
		filter := bson.M{"video_id": userFeedbacks.VideoID, "user_id": userFeedbacks.UserID}

		// 更新选项：如果文档不存在则插入
		options_ := options.Update().SetUpsert(true)

		// 执行更新操作
		_, err := userFeedbacksCollection.UpdateOne(context.TODO(), filter, update, options_)
		if err != nil {
			log.Fatal(err)
		}

	}

	// 插入文档到MongoDB
	_, err = userBehaviorCollection.InsertOne(context.TODO(), userBehavior)
	if err != nil {
		log.Fatal("插入UserBehavior文档时出错：", err)
	}

	_, err = videoAnalyticsCollection.InsertOne(context.TODO(), videoAnalytics)
	if err != nil {
		log.Fatal("插入VideoAnalytics文档时出错：", err)
	}

	_, err = seriesAnalyticsCollection.InsertOne(context.TODO(), seriesAnalytics)
	if err != nil {
		log.Fatal("插入SeriesAnalytics文档时出错：", err)
	}

	_, err = contextDataCollection.InsertOne(context.TODO(), contextData)
	if err != nil {
		log.Fatal("插入ContextData文档时出错：", err)
	}

	_, err = danmukuCollection.InsertOne(context.TODO(), danmukuData)
	if err != nil {
		log.Fatal("插入Danmuku文档时出错：", err)
	}

	log.Println("所有文档插入成功。")
}
