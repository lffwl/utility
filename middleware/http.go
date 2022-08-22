package middleware

import (
	jsoniter "github.com/json-iterator/go"
	uerror "github.com/lffwl/utility/error"
	"github.com/lffwl/utility/response"
	"io/ioutil"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	HttpOKCode                   = 200
	HttpInternalServiceErrorCode = 500
	HttpUnauthorizedCode         = 401
)

func HttpResponse(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		bytes, err := ioutil.ReadAll(r.Response.Body)
		if err != nil {
			response.Json(w, HttpInternalServiceErrorCode, err.Error())
			return
		}

		var res response.JsonResponse
		if err := json.Unmarshal(bytes, &res); err != nil {
			response.Json(w, HttpInternalServiceErrorCode, err.Error())
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
		return
	})
}
