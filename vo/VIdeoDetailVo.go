package vo

type VideoDetailVo struct {
	ID          uint       `json:"id"`
	CreatedAt   CustomTime `json:"created_at"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Url         string     `json:"ulr"`
	CoverUrl    string     `json:"cover_url"`
	Views       int        `json:"views"`
	Likes       int        `json:"likes"`
	Collections int        `json:"collections"`
	Duration    float64    `json:"duration"`
	Partition   string     `json:"partition"`
	Quality     string     `json:"quality"`
	CategoryID  uint       `json:"category_id"`
	UserID      uint       `json:"user_id"`
	Tags        string     `json:"tags"`
}
