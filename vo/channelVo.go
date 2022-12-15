package vo

type ChannelListVo struct {
	ID         uint   `json:"id"`
	Cover      string `json:"imgUrl"`
	Hot        string `json:"hot"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Vid        string `json:"videoUrl"`
	Uid        string `json:"uid"`
	Time       string `json:"time"`
	HeadImg    string `json:"headUrl"`
	Name       string `json:"name"`
	Click      string `json:"see"`
	Fans       string `json:"fans"`
	Follow     string `json:"follow"`
	Danmu      string `json:"danmu"`
	Like       string `json:"like"`
	Dislike    string `json:"dislike"`
	Collection string `json:"collection"`
	Share      string `json:"share"`
	Comment    string `json:"comment"`
}

type UpRecommendVo struct {
	Uid      string `json:"uid"`
	ImgUrl   string `json:"imgUrl"`
	VideoUrl string `json:"videoUrl"`
	Title    string `json:"title"`
}
type VideoVo struct {
	Vid string `json:"vid"`
}

type ChannelVideoInfoNumVo struct {
	Click      string `json:"see"`
	Danmu      string `json:"danmu"`
	Like       string `json:"like"`
	Dislike    string `json:"dislike"`
	Collection string `json:"collection"`
	Comment    string `json:"comment"`
	Share      string `json:"share"`
	Vid        string `json:"vid"`
	Uid        string `json:"uid"`
}
