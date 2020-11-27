package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Article struct {
	Title, Body string
	ID          int64
}

var router *mux.Router
var db *sql.DB

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		// 2. 将请求传递下去
		next.ServeHTTP(w, r)
	})
}

func getArticleByID(id string) (article Article, err error) {
	article = Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
	return
}

func articlesEditHandler(w http.ResponseWriter, r *http.Request) {
	// 获取url中的id
	id := getRouteVariable("id", r)
	// 2. 读取对应的文章数据
	article, err := getArticleByID(id)

	// 3. 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 读取成功，显示表单
		updateURL, _ := router.Get("articles.update").URL("id", id)
		data := ArticlesFormData{
			Title:  article.Title,
			Body:   article.Body,
			URL:    updateURL,
			Errors: nil,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
		logger.LogError(err)

		tmpl.Execute(w, data)
	}
}

func articlesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// 获取url中的id
	id := getRouteVariable("id", r)

	_, err := getArticleByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章不存在")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		errors := validateArticleFormData(title, body)
		if len(errors) == 0 {
			rs, err := articlesUpdateToDb(id, title, body)
			if err != nil {
				// 数据库错误
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
			}

			if n, _ := rs.RowsAffected(); n > 0 {
				showUrl, _ := router.Get("articles.show").URL("id", id)
				http.Redirect(w, r, showUrl.String(), http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何修改")
			}
		} else {
			updateUrl, _ := router.Get("articles.update").URL("id", id)
			data := ArticlesFormData{
				Title:  title,
				Body:   body,
				URL:    updateUrl,
				Errors: errors,
			}

			tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
			if err != nil {
				panic(err)
			}

			tmpl.Execute(w, data)
		}
	}
}

func (a Article) Delete() (int64, error) {
	rs, err := db.Exec("DELETE FROM articles WHERE id = " + strconv.FormatInt(a.ID, 10))

	if err != nil {
		return 0, err
	}
	// √ 更新成功，跳转到文章详情页
	if n, _ := rs.RowsAffected(); n > 0 {
		return n, nil
	}

	return 0, nil
}

func articlesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := getRouteVariable("id", r)

	article, err := getArticleByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 未找到文章")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 未出现错误，执行删除操作
		rowsAffected, err := article.Delete()

		// 4.1 发生错误
		if err != nil {
			// 应该是 SQL 报错了
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			// 4.2 未发生错误
			if rowsAffected > 0 {
				// 重定向到文章列表页
				indexURL, _ := router.Get("articles.index").URL()
				http.Redirect(w, r, indexURL.String(), http.StatusFound)
			} else {
				// Edge case
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "404 文章未找到")
			}
		}
	}
}

func getRouteVariable(parameterName string, r *http.Request) string {
	return mux.Vars(r)[parameterName]
}

func articlesUpdateToDb(id, title, body string) (rs sql.Result, err error) {
	query := "UPDATE articles SET title = ?, body = ? WHERE id = ?"
	rs, err = db.Exec(query, title, body, id)
	return
}

func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()
	route.SetRoute(router)

	router.HandleFunc("/articles/{id:[0-9]+}/edit", articlesEditHandler).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesUpdateHandler).Methods("POST").Name("articles.update")
	router.HandleFunc("/articles/{id:[0-9]+}/delete", articlesDeleteHandler).Methods("POST").Name("articles.delete")

	// 中间件：强制header内容类型为 HTML
	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
