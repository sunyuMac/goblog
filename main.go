package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog/app/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"goblog/pkg/route"
	"net/http"
)

var router *mux.Router
var db *sql.DB

func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()
	route.SetRoute(router)

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
