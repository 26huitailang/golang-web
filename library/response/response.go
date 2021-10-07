package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message interface{} `json:"message"` // 提示信息
	Details interface{} `json:"details"` // 错误详情
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func Json(c echo.Context, code ErrorCode, details interface{}, data ...interface{}) error {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	return c.JSON(
		http.StatusOK,
		JsonResponse{
			Code:    code.Code,
			Message: code.Msg,
			Details: details,
			Data:    responseData,
		})
}
