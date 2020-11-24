package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "welcome to goblog")
	} else if r.URL.Path == "/sun" {
		fmt.Fprint(w, "又来了老弟")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 not found")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
