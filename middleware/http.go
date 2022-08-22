package middleware

import (
	jsoniter "github.com/json-iterator/go"
	uerror "github.com/lffwl/utility/error"
	"github.com/lffwl/utility/response"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	HttpOKCode                   = 200
	HttpInternalServiceErrorCode = 500
	HttpUnauthorizedCode         = 401
)

func HttpResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyWriter := response.NewResponseBodyWriter(w)

		w = bodyWriter

		next.ServeHTTP(w, r)

		var res response.JsonResponse
		if err := json.Unmarshal(bodyWriter.GetBodyBytesAndReset(), &res); err != nil {
			response.Json(w, HttpInternalServiceErrorCode, err.Error())
			bodyWriter.OutPut()
			return
		}

		switch res.Code {
		case uerror.HighErrorSuccessCode:
			res.Code = HttpOKCode
		case uerror.HighErrorSqlErrorCode:
		case uerror.HighErrorParamErrorCode:
		case uerror.HighErrorFileErrorCode:
		case uerror.HighErrorServiceErrorCode:
			res.Code = HttpInternalServiceErrorCode
		case uerror.HighErrorNotAuthCode:
		case uerror.HighErrorAuthFailedCode:
			res.Code = HttpUnauthorizedCode
		}

		response.Json(w, res.Code, res.Message, res.Data)
		bodyWriter.OutPut()
		return
	})
}
