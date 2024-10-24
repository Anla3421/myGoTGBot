package dao

import (
	"database/sql"
)

var mysqlConn *sql.DB

func init() {
	CreateConn()
}

func CreateConn() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}
	mysqlConn = db
}
