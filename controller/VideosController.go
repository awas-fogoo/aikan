package controller

import (
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/services"
	"awesomeProject0511/util"
	"awesomeProject0511/vo"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type VideosController struct{}

func (c VideosController) GetCollections(ctx *gin.Context) {

	// 从 JWT 中获取用户信息
	currentUserID := util.StringToUint(ctx.Param("id"))

	// 调用服务层获取收藏列表
	videoService := services.VideoService{}
	videos, err := videoService.GetCollections(currentUserID)
	if err != nil {
		log.Println("get collections filed")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error(500, "get collections filed"))
		return
	}

	// 返回收藏列表
	ctx.JSON(http.StatusOK, dto.Success(videos))
}

func (c VideosController) GetLikes(ctx *gin.Context) {

	// 从 JWT 中获取用户信息
	currentUserID := util.StringToUint(ctx.Param("id"))

	// 调用服务层获取点赞列表
	videoService := services.VideoService{}
	users, err := videoService.GetLikes(currentUserID)
	if err != nil {
		log.Println("get likes failed")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error(500, "get likes failed"))
		return
	}

	// 返回点赞列表
	ctx.JSON(http.StatusOK, dto.Success(users))
}

func (c VideosController) GetComments(ctx *gin.Context) {
	videoID := ctx.Param("id")
	fmt.Println(videoID)
	videoService := services.VideoService{}
	res, err := videoService.GetComments(videoID)
	if err != nil {
		log.Println("get comments failed")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error(500, "get comments failed"))
		return
	}

	// 返回点赞列表
	ctx.JSON(http.StatusOK, dto.Success(res))

}

func (c VideosController) AddDanmu(ctx *gin.Context) {
	danmukuResponseVo := vo.DanmukuResponseVo{}
	err := ctx.BindJSON(&danmukuResponseVo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Error(0, "invalid request body"))
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
		ctx.JSON(http.StatusBadRequest, dto.Error(0, "one or more fields are empty"))
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

	fmt.Println(danmu)
	// 调用服务方法
	err = services.VideoService{}.AddDanmu(vid, uid, start, duration, colour, prior, content, mode, danmu.Style)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Error(0, "failed to add danmu"))
		return
	}

	ctx.JSON(http.StatusOK, dto.Success("Danmu added successfully"))
}

func (c VideosController) GetDanmu(ctx *gin.Context) {
	vid := ctx.Param("vid")
	// 调用服务方法
	danmu, err := services.VideoService{}.GetDanmu(vid)
	if err != nil {
		log.Println(err.Error())
		dto.Error(0, "get danmu false ")
		return
	}

	ctx.JSON(200, dto.Success(danmu))
}
