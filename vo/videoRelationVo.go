package vo

type VideoRelationVo struct {
	Uid       uint   `json:"uid"`
	IsLike    bool   `json:"like"`
	IsDislike bool   `json:"dislike"`
	IsCollect bool   `json:"collect"`
	IsFollow  bool   `json:"follow"`
	IsShare   bool   `json:"share"`
	Vid       string `json:"vid"`
}
