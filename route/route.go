package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Router = mux.NewRouter()

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		// checkError(err)
		return ""
	}

	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	return mux.Vars(r)[parameterName]
}