package controller

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	// 连接数据库
	db := common.InitDB()
	defer db.Close()

	//查询上一个uid是多少进行递增
	user := model.User{}
	db.AutoMigrate(&user)
	db.Last(&user)

	// ajax获取参数
	c.Bind(&user)
	name := user.Name
	email := user.Email
	password := user.Password
	atoi_uid, _ := strconv.Atoi(user.Uid)
	atoi_uid += 1

	/*  postman 调试参数
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	*/

	// 数据验证 如果名称没有传，则自动生成一个10位的字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if len(email) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "邮件的地址格式错误"})
		return
	}

	//判断邮箱是否存在
	if isEmailExist(db, email) {
		c.JSON(422, gin.H{"code": 422, "msg": "邮箱已经存在"})
		log.Println(1)
		return
	}

	if len(password) < 6 {
		c.JSON(422, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

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
	DB := common.InitDB()
	email := requestUser.Email
	password := requestUser.Password

	log.Println(email, password)
	// 数据验证
	if len(email) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "邮件的地址格式错误"})
		return
	}
	if len(password) < 6 {
		c.JSON(422, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

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
