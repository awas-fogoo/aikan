package controller

import (
	"github.com/gin-gonic/gin"
)

func AutoCreateUser(c *gin.Context) {

	//addUser(c)

}

//func addUser(c *gin.Context) {
//	db := common.DB
//	var users []*model.User
//	err := util.WithTransaction(db, func(tx *gorm.DB) error {
//		// 增删改操作
//		for i := 0; i < 20; i++ {
//			user := &model.User{
//				Username:  fmt.Sprintf("user%d", i),
//				Nickname:  fmt.Sprintf("nickname%d", i),
//				Email:     fmt.Sprintf("user%d@bup.pub", i),
//				AvatarUrl: fmt.Sprintf("https://bup.pub/%d.jpg", i),
//				Gender:    "男",
//				Age:       uint(i),
//			}
//			user.Auth.Password = fmt.Sprintf("$2a$10$JfgQTe0bp.Xgd3zmYcHOwObCQ2nC6eMNIppT2z/jwdQonFZbs4rnK")
//			if err := tx.Create(user).Error; err != nil {
//				return err
//			}
//			users = append(users, user)
//		}
//		return nil
//	})
//	if err != nil {
//		// 发生错误，需要回滚事务
//		log.Fatalf(fmt.Sprintf("Failed to execute transactions: %v", err))
//	}
//	c.JSON(200, dto.Success(users))
//	return
//}
