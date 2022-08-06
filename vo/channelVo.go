package vo

type ChannelListVo struct {
	CreateTime string `json:"time"`
	ImgUrl     string `json:"imgUrl"`
	Hot        string `json:"hot"`
	Title      string `json:"title"`
	HeadUrl    string `json:"headUrl"`
	Name       string `json:"name"`
	Fans       string `json:"fans"`
	Follow     string `json:"follow"`
	See        string `json:"see"`
	Danmu      string `json:"danmu"`
	VideoUrl   string `json:"videoUrl"`
	Desc       string `json:"desc"`
	Like       string `json:"like"`
	Dislike    string `json:"dislike"`
	Collection string `json:"collection"`
	Share      string `json:"share"`
	Comment    string `json:"comment"`
	Uid        string `json:"uid"`
}

type UpRecommendVo struct {
	Uid      string `json:"uid"`
	ImgUrl   string `json:"imgUrl"`
	VideoUrl string `json:"videoUrl"`
	Title    string `json:"title"`
}
