package vo

type UserVo struct {
	Name     string
	Email    string
	Password string
	Code     string
}

type UserInfoVo struct {
	Fans    string `json:"fans"`
	Name    string `json:"name"`
	HeadImg string `json:"headUrl"`
}
