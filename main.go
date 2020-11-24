package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "welcome to goblog")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 not found")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "又来了老弟")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)

	router.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		write := ""
		switch r.Method {
		case "POST":
			write = "创建新文章"
			break
		case "GET":
			write = "获取文章信息"
			break
		default:
			write = "请求方式错误"
		}

		fmt.Fprint(w, write)
	})

	http.ListenAndServe(":3000", router)
}
