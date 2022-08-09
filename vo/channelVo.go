package vo

import "time"

type ChannelListVo struct {
	ID        uint      `json:"id"`
	Cover     string    `json:"imgUrl"`
	Hot       string    `json:"hot"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Vid       string    `json:"videoUrl"`
	Uid       string    `json:"uid"`
	CreatedAt time.Time `json:"time"`

	//HeadUrl    string `json:"headUrl"`
	//Name       string `json:"name"`
	//Fans       string `json:"fans"`
	//Follow     string `json:"follow"`
	//Danmu      string `json:"danmu"`
	//Like       string `json:"like"`
	//Dislike    string `json:"dislike"`
	//Collection string `json:"collection"`
	//Share      string `json:"share"`
	//Comment    string `json:"comment"`
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
