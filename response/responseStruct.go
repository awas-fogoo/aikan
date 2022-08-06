package response

import "github.com/gin-gonic/gin"

type ResStruct struct {
	Ret  bool  `json:"ret"`
	Data gin.H `json:"data"` //数据
}
