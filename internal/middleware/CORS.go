package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"127.0.0.1"}
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
