package middleware

import (
	"keydol/internal/consts"
	"net/http"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

type DefaultResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

// CORS 跨域设置
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// HandleResponse 统一处理响应，仿照ghttp.MiddlewareHandlerResponse设计
func HandleResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg      string
		res, err = r.GetHandlerResponse()
		code     = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			code = gcode.CodeNotFound
		case http.StatusForbidden:
			code = gcode.CodeNotAuthorized
		default:
			code = gcode.CodeUnknown
		}
	} else {
		code = gcode.New(consts.Success, "成功", nil)
		msg = "成功"
	}

	defaultRes := DefaultResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
		Success: code.Code() == consts.Success,
	}

	responseString := gjson.New(defaultRes).MustToJsonString()

	// 返回
	r.Response.WriteJson(responseString)

	// 暂时先决定WriteJson后退出路由
	r.Exit()
}
