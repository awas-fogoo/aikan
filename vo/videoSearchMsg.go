package vo

type VideoMsg struct {
	Title            string  `gorm:"size:255"`
	Description      *string `gorm:"type:text"`
	Uploader         *string `gorm:"size:255"`
	Duration         int
	StoryId          int
	Category         *string `gorm:"size:100"`
	Resolution       *string `gorm:"size:100"`
	BelongsToSeries  *uint   `gorm:"index;constraint:OnDelete:SET NULL"`
	CoverImageUrl    *string `gorm:"default:null"`
	CollectionNumber int
}
