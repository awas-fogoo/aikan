package controller

//
//import (
//	"awesomeProject0511/common"
//	"awesomeProject0511/dto"
//	"awesomeProject0511/model"
//	"awesomeProject0511/server"
//	"awesomeProject0511/util"
//	"errors"
//	"github.com/gin-gonic/gin"
//	"github.com/jinzhu/gorm"
//	"golang.org/x/crypto/bcrypt"
//	"log"
//	"regexp"
//	"time"
//)
//
//func SendVerificationCode(c *gin.Context) {
//	email := c.PostForm("email")
//	//ajax获取参数
//	//regUser := vo.UserVo{}
//	//c.Bind(&regUser)
//	//email := regUser.Email
//	if !isEmailValid(email) {
//		c.JSON(0, dto.Error(0, "邮件的地址格式错误"))
//		return
//	}
//
//	// 连接数据库
//	db := common.InitDB()
//	defer db.Close()
//
//	//判断邮箱是否存在
//	existEmail := util.IsFieldExist(db, "email", email)
//	if existEmail {
//		c.JSON(0, dto.Error(0, "邮箱已经存在"))
//		return
//	}
//
//	rdb := common.InitCache()
//	defer rdb.Close()
//	ctx := common.Ctx
//	randomCode := util.RandomCode(6)
//	key := util.ReEmail(email)
//	res, _ := rdb.Get(ctx, key).Result()
//	if res != "" {
//		duration, _ := rdb.TTL(ctx, key).Result()
//		if duration >= 240000000000 {
//			c.JSON(0, dto.Error(0, "请务频繁发送"))
//			return
//		}
//		return
//	}
//
//	// 创建redis三分钟验证码有效期
//	err := rdb.Set(ctx, key, randomCode, time.Second*300).Err()
//	if err != nil {
//		c.JSON(0, dto.Error(0, "验证码有效期错误"))
//		return
//	}
//
//	// 发送验证码到这个邮箱
//	server.SendVerificationCode(email, randomCode)
//	c.JSON(0, dto.Error(0, "验证码发送成功"))
//}
//
//func Register(c *gin.Context) {
//
//	// ajax获取参数
//	//regUser := vo.UserVo{}
//	//c.Bind(&regUser)
//
//	/****** postman 调试参数	******/
//	email := c.PostForm("email")
//	//email := regUser.Email
//	if !isEmailValid(email) {
//		c.JSON(0, dto.Error(0, "邮件的地址格式错误"))
//		return
//	}
//
//	/****** postman 调试参数	******/
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//	nickname := c.PostForm("nickname")
//	if len(username) == 0 {
//		c.JSON(0, dto.Error(0, "用户名不能为空，且唯一"))
//		return
//	}
//	// 连接数据库
//	db := common.InitDB()
//	defer db.Close()
//
//	existUsername := util.IsFieldExist(db, "username", username)
//	if existUsername {
//		c.JSON(0, dto.Error(0, "用户名已经存在"))
//		return
//	}
//
//	//name := regUser.Name
//	//password := regUser.Password
//
//	// 数据验证 如果名称没有传，则自动生成一个10位的字符串
//	if len(nickname) == 0 {
//		nickname = util.RandomString(10)
//	}
//
//	if len(password) < 6 && len(password) == 0 {
//		c.JSON(0, dto.Error(0, "密码不能少于6位"))
//		return
//	}
//
//	/****** postman 调试参数	******/
//	code := c.PostForm("code")
//	//randomCode := regUser.Code
//
//	// 校验验证码是否正确
//	if len(code) == 0 {
//		c.JSON(0, dto.Error(0, "请发送验证码"))
//		return
//	}
//	rdb := common.InitCache()
//	ctx := common.Ctx
//
//	// redis验证码核对
//	regEmail := util.ReEmail(email)
//	rdbCode, _ := rdb.Get(ctx, regEmail).Result()
//	if rdbCode == "" || rdbCode != code {
//		c.JSON(0, dto.Error(0, "验证码错误"))
//		return
//	}
//
//	//加密处理
//	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		c.JSON(500, gin.H{})
//	}
//	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
//
//	log.Println(username, email, password)
//
//	//创建用户
//	var newUser = model.User{
//		Username: username,
//		Email:    email,
//		Password: encodePWD,
//		Nickname: nickname,
//	}
//	db.Create(&newUser)
//
//	// 发送token
//	token, err := common.ReleaseToken(newUser)
//	if err != nil {
//		log.Printf("token generate error : %v", err)
//		c.JSON(0, dto.Error(0, "系统异常"))
//
//	}
//	rdb.Del(ctx, regEmail) // 删除key
//	//返回结果
//	c.JSON(200, dto.Success(token))
//	return
//}
//
//func Login(c *gin.Context) {
//	// ajax获取参数
//	//requestUser := model.User{}
//	//c.Bind(&requestUser)
//	//email := requestUser.Email
//	//password := requestUser.Password
//
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//	log.Println("login:", username, password)
//	// 数据验证
//	if len(username) == 0 {
//		c.JSON(0, dto.RetDTO{Message: "请输入用户名"})
//		return
//	}
//	if len(password) < 6 && len(password) == 0 {
//		c.JSON(0, dto.RetDTO{Message: "请输入密码,密码长度不能低于6位"})
//		return
//	}
//
//	db := common.InitDB()
//	// 判断username是否存在
//	user, err := verityPwd(db, username, password)
//	if err != nil {
//		log.Println(err)
//		c.JSON(0, dto.RetDTO{Message: "用户名或密码错误"})
//	} else {
//		// 发送token
//		token, err := common.ReleaseToken(*user)
//		if err != nil {
//			c.JSON(500, dto.RetDTO{Message: "系统异常"})
//			log.Printf("token generate error : %v", err)
//		}
//		c.JSON(0, dto.RetDTO{Message: "login success", Data: token})
//	}
//}
//
//var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
//
//func isEmailValid(e string) bool {
//	if len(e) < 3 && len(e) > 254 {
//		return false
//	}
//	return emailRegex.MatchString(e)
//}
//
//func verityPwd(db *gorm.DB, username string, password string) (*model.User, error) {
//	var user model.User
//	result := db.Where("username = ? OR email = ?", username, username).First(&user)
//	if result.RowsAffected == 0 {
//		// 用户不存在
//		return nil, errors.New("user not found")
//	} else if result.Error != nil {
//		// 查询出错，处理错误
//		return nil, result.Error
//	} else if !CheckPasswordHash(password, user.Password) {
//		// 密码不匹配
//		return nil, errors.New("invalid password")
//	} else {
//		// 登录成功
//		return &user, nil
//	}
//}
//func CheckPasswordHash(password string, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}
//
//func SearchUserController(c *gin.Context) {
//	server.SearchUserServer(c)
//}
