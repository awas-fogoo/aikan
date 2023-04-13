package dto

type RetDTO struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewRetDTO(code int, message string, data interface{}) *RetDTO {
	return &RetDTO{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// Success 返回成功结果
func Success(data interface{}) *RetDTO {
	return NewRetDTO(0, "success", data)
}

// Error 返回错误结果
func Error(code int, message string) *RetDTO {
	return NewRetDTO(code, message, nil)
}

/* e.g.

func GetUserInfo(userId uint) *RetDTO {
    user := &model.User{}
    if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
        return Error(1001, fmt.Sprintf("Failed to get user: %v", err))
    }
    return Success(user)
}
在上述示例代码中，我们通过 Error() 方法创建了一个返回错误结果对象，并在错误信息中输出了详细的错误信息。
在返回成功结果时，我们则可以使用 Success() 方法创建一个新的返回对象，并将获取到的用户信息作为返回数据。
这样做可以有效地管理和输出返回结果信息，提高系统的可读性和可维护性。


2000 OK：服务器成功返回请求数据。
2010 Created：服务器已经成功创建了资源。
2040 No Content：没有返回任何内容。
4000 Bad Request：请求参数不正确。
4010 Unauthorized：访问未经授权。
4030 Forbidden：禁止访问。
4040 Not Found：请求的资源不存在。
4060 Not Acceptable：请求的格式不可得。
4100 Gone：请求的资源已经不存在。
4220 Unprocessable Entity：输入验证错误。
5000 Internal Server Error：服务器内部错误。


*/
