package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"one/common"
	"one/dto"
	"one/model"
	"one/services"
	"one/util"
	"regexp"
	"time"
)

func SendVerificationCode(c *gin.Context) {
	var data struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(200, dto.Error(4001, "Invalidated request data"))
		return
	}
	email := data.Email
	if !isEmailValid(email) {
		c.JSON(200, dto.Error(4002, "Incorrectly formatted email address"))
		return
	}

	// 连接数据库
	db := common.DB

	//判断邮箱是否存在
	existEmail := IsFieldExist(db, "email", email)
	if existEmail {
		c.JSON(200, dto.Error(4003, "Email already exists"))
		return
	}

	rdb := common.RDB
	ctx := context.TODO()
	randomCode := util.RandomCode(6)
	key := util.ReEmail(email)

	res, _ := rdb.Get(ctx, key).Result()
	if res != "" {
		duration, _ := rdb.TTL(ctx, key).Result()
		if duration >= 240000000000 {
			c.JSON(200, dto.Error(4004, "Place do not send frequently"))
			return
		}
		return
	}

	// 创建redis三分钟验证码有效期
	err := rdb.Set(ctx, key, randomCode, time.Minute*5).Err()
	if err != nil {
		c.JSON(200, dto.Error(5001, "Failed to set the expiration time for the verification code"))
		return
	}

	// 发送验证码到这个邮箱
	services.SendVerificationCode(email, randomCode)
	c.JSON(200, dto.Error(5002, "Failed to send the verification code"))
}

// Register 定义注册函数
func Register(c *gin.Context) {
	var data struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(200, dto.Error(4001, "Invalid request data"))
		return
	}
	// 验证用户名长度
	if len(data.Username) < 3 || len(data.Username) > 10 {
		c.JSON(200, dto.Error(4002, "Username length should be between 3 and 10 characters"))
		return
	}

	// 验证密码长度
	if len(data.Password) < 6 || len(data.Password) > 20 {
		c.JSON(200, dto.Error(4003, "Password length should be between 6 and 20 characters"))
		return
	}
	var existingUser model.User
	if err := common.DB.Where("username = ? OR email = ?", data.Username, data.Email).First(&existingUser).Error; err != gorm.ErrRecordNotFound {
		c.JSON(200, dto.Error(4004, "Username or email already exists"))
		return
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		c.JSON(200, dto.Error(5001, "Password encryption failed"))
		return
	}

	newUser := model.User{
		Username: data.Username,
		Email:    data.Email,
		Password: string(hashedPwd),
	}

	tx := common.DB.Begin()
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		c.JSON(200, dto.Error(5002, "Failed to create user"))
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err)
		c.JSON(200, dto.Error(5003, "User transaction commit failed"))
		return
	}

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		log.Printf("token generate error: %v", err)
		c.JSON(200, dto.Error(5004, "System error"))
		return
	}

	c.JSON(200, dto.RetDTO{
		Message: "Creation successful",
		Data: gin.H{
			"token": token,
			"userinfo": gin.H{
				"id":       newUser.ID,
				"username": newUser.Username,
				"email":    newUser.Email,
			},
		},
	})
}
func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(200, dto.Error(4001, "Invalid request payload"))
		return
	}

	username := loginData.Username
	password := loginData.Password
	log.Println("login attempt:", username)

	// 这里添加密码验证逻辑
	tx := common.DB.Begin()
	user, err := verifyPwd(tx, username, password)
	if err != nil {
		log.Println(err)
		c.JSON(200, dto.Error(4002, "Username or password incorrect"))
		return
	}

	// 生成Token
	token, err := common.ReleaseToken(*user)
	if err != nil {
		c.JSON(200, dto.Error(5001, "System error"))
		log.Printf("token generate error: %v", err)
		return
	}

	// 构建响应信息
	userInfo := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"nickname": user.Nickname,
	}
	info := struct {
		Token string      `json:"token"`
		User  interface{} `json:"userinfo"`
	}{
		Token: token,
		User:  userInfo,
	}

	c.JSON(200, dto.RetDTO{Message: "Login successful", Data: info})
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func verifyPwd(db *gorm.DB, username string, password string) (*model.User, error) {
	var user model.User
	// 用用户名或者邮箱来查找用户
	result := db.Where("username = ? OR email = ?", username, username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error while querying for user: %v", result.Error)
	}

	// 验证密码是否正确
	if !CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func IsFieldExist(db *gorm.DB, field string, value string) bool {
	var user model.User
	result := db.Where(field+" = ?", value).First(&user)
	if result.RowsAffected == 0 {
		// 未查询到记录，说明字段不存在
		return false
	} else if result.Error != nil {
		// 查询出错，处理错误
		log.Fatalln(result.Error)
	} else {
		// 查询到记录，说明字段已存在
		return true
	}
	return false
}
