package route

import (
	"github.com/gorilla/mux"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"net/http"
)

var route *mux.Router

// SetRoute 设置路由实例，以供 Name2URL 等函数使用
func SetRoute(r *mux.Router) {
	route = r
}

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}

	return config.GetString("app.url") + url.String()
}

// GetRouteVariable 获取路由中的参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	return mux.Vars(r)[parameterName]
}
