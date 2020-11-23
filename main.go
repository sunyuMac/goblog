package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "anyone")
	} else if r.URL.Path == "/sun" {
		fmt.Fprint(w, "又来了老弟")
	} else {
		fmt.Fprint(w, "404 not found")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
