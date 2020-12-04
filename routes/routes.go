package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

// RegisterRoutes 注册路由
func RegisterRoutes(r *mux.Router) {
	//静态资源
	registerStaticRoutes(r)
	// web 路由
	registerWebRoutes(r)
}

func registerStaticRoutes(r *mux.Router) {
	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))
}
