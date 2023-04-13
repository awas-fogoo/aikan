package vo

type UserVo struct {
	ID        uint   `json:"id"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatar_url"`
}
