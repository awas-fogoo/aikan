package vo

type SearchVideoVo struct {
	ID          uint       `json:"id"`
	CreatedAt   CustomTime `json:"created_at"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Url         string     `json:"url"`
	CoverUrl    string     `json:"cover_url"`
	Duration    float64    `json:"duration"`
	CategoryID  uint       `json:"category_id"`
	UserID      uint       `json:"user_id"`
	Tags        string     `json:"tag_name"`
}
