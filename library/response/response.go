package response

import (
	"github.com/labstack/echo"
	"net/http"
)

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func Json(c echo.Context, code int, message string, data ...interface{}) error {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	return c.JSON(
		http.StatusOK,
		JsonResponse{
			Code:    code,
			Message: message,
			Data:    responseData,
		})
}
