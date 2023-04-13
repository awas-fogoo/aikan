package util

import (
	"awesomeProject0511/model"
	"github.com/jinzhu/gorm"
)

func IsFieldExist(db *gorm.DB, field string, value string) bool {
	var user model.User
	result := db.Where(field+" = ?", value).First(&user)
	if result.RowsAffected == 0 {
		// 未查询到记录，说明字段不存在
		return false
	} else if result.Error != nil {
		// 查询出错，处理错误
		panic(result.Error)
	} else {
		// 查询到记录，说明字段已存在
		return true
	}
}
