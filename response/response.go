package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

// Continue 100  继续，服务器收到请求，需要请求者继续执行操作
func Continue(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusContinue, 100, data, msg)
}

// Success 200 成功，操作被成功接收并处理
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

// MovedPermanently 301 永久移动。请求的资源已被永久的移动到新URI，返回信息会包括新的URI，浏览器会自动定向到新URI。今后任何新的请求都应使用新的URI代替
func MovedPermanently(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusMovedPermanently, 301, data, msg)
}

// BadRequest 400 客户端错误，请求包含语法错误或无法完成请求
func BadRequest(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}

// Unauthorized 401	请求要求用户的身份认证
func Unauthorized(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusUnauthorized, 401, data, msg)
}

// 	Forbidden 403 服务器理解请求客户端的请求，但是拒绝执行此请求
func Forbidden(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusForbidden, 403, data, msg)
}

// InternalServerError 500 服务器内部错误，无法完成请求
func InternalServerError(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusInternalServerError, 500, data, msg)
}

// NotImplemented 501 服务器不支持请求的功能，无法完成请求
func NotImplemented(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusNotImplemented, 501, data, msg)
}
