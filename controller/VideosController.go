package controller

import (
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/services"
	"awesomeProject0511/util"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
	"log"
)

type VideosController struct{}

func (VideosController) GetCollections(ctx *gin.Context) {

	// 从 JWT 中获取用户信息
	currentUserID := util.StringToUint(ctx.Param("id"))

	// 调用服务层获取收藏列表
	videoService := services.VideoService{}
	videos, err := videoService.GetCollections(currentUserID)
	if err != nil {
		log.Println("get collections filed")
		ctx.AbortWithStatusJSON(200, dto.Error(5000, "Get collections filed"))
		return
	}

	// 返回收藏列表
	ctx.JSON(200, dto.Success(videos))
}

func (VideosController) GetLikes(ctx *gin.Context) {

	// 从 JWT 中获取用户信息
	currentUserID := util.StringToUint(ctx.Param("id"))

	// 调用服务层获取点赞列表
	videoService := services.VideoService{}
	users, err := videoService.GetLikes(currentUserID)
	if err != nil {
		log.Println("get likes failed")
		ctx.AbortWithStatusJSON(200, dto.Error(5000, "Get likes failed"))
		return
	}

	// 返回点赞列表
	ctx.JSON(200, dto.Success(users))
}

func (VideosController) GetComments(ctx *gin.Context) {
	videoID := ctx.Param("id")
	videoService := services.VideoService{}
	result, err := videoService.GetComments(videoID)
	if err != nil {
		ctx.AbortWithStatusJSON(200, dto.Error(5000, "Get comments failed"))
		return
	}

	// 返回点赞列表
	ctx.JSON(200, dto.Success(result))

}

func (VideosController) AddDanmu(ctx *gin.Context) {
	danmukuResponseVo := vo.DanmukuResponseVo{}
	err := ctx.BindJSON(&danmukuResponseVo)
	if err != nil {
		ctx.JSON(200, dto.Error(4000, "Invalid request body"))
		return
	}

	vid := util.StringToUint(ctx.Param("vid"))
	//uid := danmukuResponseVo.UserID
	uid := uint(999999)
	start := danmukuResponseVo.Start
	duration := danmukuResponseVo.Duration
	prior := danmukuResponseVo.Prior
	colour := danmukuResponseVo.Colour
	content := danmukuResponseVo.Content
	mode := danmukuResponseVo.Mode
	style := danmukuResponseVo.Style

	// 检查多个变量是否为空
	if vid == 0 || uid == 0 || start == 0 || duration == 0 || len(content) == 0 {
		ctx.JSON(200, dto.Error(4000, "One or more fields are empty"))
		return
	}

	// 创建弹幕对象
	danmu := model.Danmuku{
		VideoID:  vid,
		UserID:   uid,
		Content:  content,
		Start:    start,
		Duration: duration,
		Prior:    prior,
		Colour:   colour,
		Mode:     mode,
		Style: model.DanmukuStyle{
			Color:           style.Color,
			FontSize:        style.FontSize,
			Border:          style.Border,
			BorderRadius:    style.BorderRadius,
			Padding:         style.Padding,
			BackgroundColor: style.BackgroundColor,
		},
	}

	// 调用服务方法
	videoService := services.VideoService{}
	err = videoService.AddDanmu(vid, uid, start, duration, colour, prior, content, mode, danmu.Style)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(200, dto.Error(5000, "Failed to add danmu"))
		return
	}

	ctx.JSON(200, dto.Success(nil))
}

func (VideosController) GetDanmus(ctx *gin.Context) {
	vid := ctx.Param("vid")
	// 调用服务方法
	danmu, err := services.VideoService{}.GetDanmu(vid)
	if err != nil {
		log.Println(err.Error())
		dto.Error(5000, "Get danmu failed ")
		return
	}

	ctx.JSON(200, dto.Success(danmu))
}

func (VideosController) UploadVideo(c *gin.Context) {
	//title := c.PostForm("title")
	//desc := c.PostForm("description")
	//url := c.PostForm("url")
	//coverUrl := c.PostForm("cover_url")
	//tags := c.PostForm("tags")
	//categoryId := util.StringToUint(c.PostForm("category_id"))
	uploadVideoVo := vo.SearchVideoVo{}
	if err := c.Bind(&uploadVideoVo); err != nil {
		c.JSON(200, dto.Error(4000, "Invalid request payload"))
		return
	}
	getUser, _ := c.Get("user")
	userDto := dto.ToUserDTO(getUser.(model.User))
	// 调用服务层上传视频函数
	videoService := services.VideoService{}
	videoID, err := videoService.UploadVideoServer(uploadVideoVo.Title, uploadVideoVo.Description, uploadVideoVo.Url, uploadVideoVo.CoverUrl, uploadVideoVo.Tags, uploadVideoVo.CategoryID, userDto.ID)
	if err != nil {
		log.Println(err)
		c.JSON(200, dto.Error(err.Code, err.Message))
		return
	}
	c.JSON(200, dto.Success(map[string]interface{}{
		"vid": videoID, // 请将 videoID 替换为实际的视频 ID
	}))
}

func (VideosController) AddLikeVideo(c *gin.Context) {
	user, _ := c.Get("user")
	userDto := dto.ToUserDTO(user.(model.User))
	videoID := util.StringToUint(c.Param("id"))
	userID := userDto.ID

	if videoID == 0 {
		c.JSON(200, dto.Error(4000, "Wrong Video ID"))
	}
	videoService := services.VideoService{}
	response := videoService.AddLikeVideo(videoID, userID)
	c.JSON(200, response)
}

/*
通常来说，控制层应该负责处理 HTTP 请求、验证输入、构建响应以及将请求传递给服务层来执行具体的业务逻辑。
服务层负责实际的业务逻辑处理，比如视频上传、标签处理等。
*/

func (VideosController) GetVideoList(c *gin.Context) {
	services.GetVideoListServer(c)
}

func (VideosController) GetVideoDetail(c *gin.Context) {
	services.GetVideoDetailServer(c)
}

func (VideosController) SearchVideo(c *gin.Context) {
	services.SearchVideoServer(c)
}

func (VideosController) AddCollectionVideo(c *gin.Context) {
	services.AddCollectionVideoServer(c)
}

func (VideosController) AddCommentVideo(c *gin.Context) {
	services.AddCommentVideoServer(c)
}

func (VideosController) GetCommentVideo(c *gin.Context) {
	services.GerCommentVideoServer(c)
}

func (VideosController) GetHotVideo(c *gin.Context) {
	services.GetHotVideoServer(c)
}

func (VideosController) VideoAddressForward(c *gin.Context) {
	services.VideoStreamServer(c)
}
