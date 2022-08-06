package controller

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/server"
	"awesomeProject0511/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func SendCode(c *gin.Context) {
	//email := c.PostForm("email")
	// ajax获取参数
	regUser := model.RegUser{}
	c.Bind(&regUser)
	email := regUser.Email
	fmt.Println(email)
	if !isEmailValid(email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "邮件的地址格式错误"})
		return
	}

	// 连接数据库
	db := common.InitDB()
	defer db.Close()

	//判断邮箱是否存在
	if isEmailExist(db, email) {
		c.JSON(422, gin.H{"code": 422, "msg": "邮箱已经存在"})
		return
	}
	rdb := common.InitCache()
	defer rdb.Close()
	ctx := common.Ctx
	randomCode := util.RandomCode(6)
	key := util.ReEmail(email)
	res, _ := rdb.Get(ctx, key).Result()
	if res != "" {
		duration, _ := rdb.TTL(ctx, key).Result()
		if duration >= 240000000000 {
			c.JSON(301, gin.H{
				"msg": "请务频繁发送",
			})
			return
		}
		c.JSON(301, gin.H{
			"msg": "邮件错误",
		})
		return
	}

	// 创建redis三分钟验证码有效期
	err := rdb.Set(ctx, key, randomCode, time.Second*300).Err()
	if err != nil {
		fmt.Println("err")
		return
	}

	// 发送验证码到这个邮箱
	server.SendVerificationCode(email, randomCode)
	c.JSON(200, gin.H{
		"msg": "验证码发送成功",
	})

}

func Register(c *gin.Context) {

	// ajax获取参数
	regUser := model.RegUser{}
	c.Bind(&regUser)

	/****** postman 调试参数	******/
	//email := c.PostForm("email")
	email := regUser.Email
	if !isEmailValid(email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "邮件的地址格式错误"})
		return
	}

	/****** postman 调试参数	******/
	//name := c.PostForm("name")
	//password := c.PostForm("password")
	name := regUser.Name
	password := regUser.Password

	// 数据验证 如果名称没有传，则自动生成一个10位的字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	if len(password) < 6 && len(password) == 0 {
		c.JSON(422, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	/****** postman 调试参数	******/
	//randomCode := c.PostForm("code")
	randomCode := regUser.Code

	// 校验验证码是否正确
	if len(randomCode) == 0 {
		c.JSON(422, gin.H{"code": 422, "msg": "请发送验证码"})
		return
	}
	rdb := common.InitCache()
	ctx := common.Ctx

	// 创建redis-key
	regEmail := util.ReEmail(email)
	rdbCode, _ := rdb.Get(ctx, regEmail).Result()
	if rdbCode == "" || rdbCode != randomCode {
		c.JSON(422, gin.H{
			"msg": "校验错误，请重新发送",
		})
		return
	}

	// 连接数据库
	db := common.InitDB()
	defer db.Close()

	//查询上一个uid是多少进行递增
	user := model.User{}
	db.Last(&user)
	atoi_uid, _ := strconv.Atoi(user.Uid)
	atoi_uid += 1

	//加密处理
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{})
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可

	log.Println(name, email, password)

	//创建用户
	var newUser = model.User{
		Name:     name,
		Email:    email,
		Password: encodePWD,
		Uid:      strconv.Itoa(atoi_uid),
	}
	db.Create(&newUser)

	// 发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "系统异常"})
		log.Printf("token generate error : %v", err)
	}
	rdb.Del(ctx, regEmail) // 删除key
	//返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"ok":   "创建成功",
	})
}

func Login(c *gin.Context) {
	// ajax获取参数
	requestUser := model.User{}
	c.Bind(&requestUser)
	email := requestUser.Email
	password := requestUser.Password

	log.Println(email, password)
	// 数据验证
	if !isEmailValid(email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "邮件的地址格式错误"})
		return
	}
	if len(password) < 6 && len(password) == 0 {
		c.JSON(422, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	DB := common.InitDB()
	// 判断email是否存在
	var user model.User
	DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 发送token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "系统异常"})
		log.Printf("token generate error : %v", err)
	}

	//返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"ok":   "登入成功",
	})
}

// user.(model.User) 类型断言 可以说成是类型强制转换
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": dto.ToUserDto(user.(model.User))})
}

func isEmailExist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
