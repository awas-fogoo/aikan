package vo

type DetailMsg struct {
	Id            int
	Title         string  `gorm:"size:255"`
	CoverImageUrl *string `gorm:"default:null"` // 封面地址
	TotalEpisodes int     `gorm:"default:1"`    // 总集数
}
