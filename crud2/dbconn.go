package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// the global variable in the dops package
var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goblog"

	//dbConn, err = sql.Open("mysql", "root:my-pass@tcp(localhost:3306)/video_server?charset=utf8")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
}
