package response

import (
	uerror "192.168.0.209/wl/utility/error"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

const MessageSuccess = "OK"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// JsonResponse 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"result"`  // 返回数据(业务接口定义具体数据结构)
}

// Json 标准返回结果数据结构封装。
func Json(w http.ResponseWriter, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	w.Header().Set("content-type", "text/json")
	res, _ := json.Marshal(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
	_, _ = w.Write(res)
}

func HttpErrorJson(w http.ResponseWriter, highError uerror.HighError, data ...interface{}) {
	Json(w, highError.Code, highError.Error.Error(), data)
}

func HttpSuccessJson(w http.ResponseWriter, message string, data ...interface{}) {
	if message == "" {
		message = MessageSuccess
	}
	Json(w, uerror.HighErrorSuccessCode, message, data)
}
