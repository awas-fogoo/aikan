# 视频完播率
```golang
要记录视频的完播率，我们需要在视频播放期间收集。具体而言，我们可以记录以下数据：

视频播放开始时间
视频总时长
视频播放结束时间
每个视频段的开始时间和结束时间（如果有的话）
每个视频段的播放时间
在收集了这些数据后，我们就可以计算出视频的完播率，即所有视频段播放时间总和与视频总时长之间的比率。例如，如果视频总共有10个片段，每个片段的播放时间为60秒，总播放时间为600秒，而视频总时长为1200秒，则该视频的完播率为600 / 1200 = 0.5。

以下是一个示例代码，用于记录视频的完播率：

type VideoPlaybackLog struct {
    VideoID        uint
    UserID         uint
    StartTime      time.Time
    EndTime        time.Time
    TotalPlayTime  float64
    SegmentPlayLog []SegmentPlayLog
}

type SegmentPlayLog struct {
    SegmentID    uint
    StartTime   time.Time
    EndTime     time.Time
    PlayTime    float64
}

// 开始播放视频
func StartVideoPlayback(db *gorm.DB, videoID, userID uint) (*VideoPlaybackLog, error) {
    log := &VideoPlaybackLog{
        VideoID:   videoID,
        UserID:    userID,
        StartTime: time.Now(),
    }

    err := db.Create(log).Error
    if err != nil {
        return nil, err
    }

    return log, nil
}

// 记录视频片段播放情况
func LogSegmentPlay(db *gorm.DB, logID, segmentID uint, startTime, endTime time.Time) error {
    var log VideoPlaybackLog
    err := db.Where("id = ?", logID).Preload("SegmentPlayLog").First(&log).Error
    if err != nil {
        return err
    }

    playTime := endTime.Sub(startTime).Seconds()

    segmentLog := &SegmentPlayLog{
        SegmentID:  segmentID,
        StartTime:  startTime,
        EndTime:    endTime,
        PlayTime:   playTime,
    }

    log.SegmentPlayLog = append(log.SegmentPlayLog, *segmentLog)
    log.TotalPlayTime += playTime

    return db.Save(&log).Error
}

// 结束视频播放
func EndVideoPlayback(db *gorm.DB, logID uint) error {
    var log VideoPlaybackLog
    err := db.Where("id = ?", logID).Preload("SegmentPlayLog").First(&log).Error
    if err != nil {
        return err
    }

    log.EndTime = time.Now()

    return db.Save(&log).Error
}
在这个示例中，我们定义了两个模型，VideoPlaybackLog和SegmentPlayLog，分别用于记录视频播放日志以及每个视频片段的播放情况。我们还编写了三个函数，StartVideoPlayback，LogSegmentPlay和EndVideoPlayback，用于开始视频播放，记录视频片段播放情况以及结束视频播放。

例如，当用户开始播放视频时，我们可以调用StartVideoPlayback函数，并记录其开始时间。当视频播放过程中用户观看了一个新片段时，我们可以调用LogSegmentPlay函数，并记录该片段的开始时间和结束时间。最后，当用户完成视频观看并退出时，我们可以调用EndVideoPlayback函数并记录视频的结束时间。

这样，我们就可以在数据库中检索相关日志记录，并计算所有视频段播放时间总和，以及视频总时长之间的比率，从而得出视频的完播率。
```