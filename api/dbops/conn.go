package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:root123456@tcp(127.0.0.1:3306)/video?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
